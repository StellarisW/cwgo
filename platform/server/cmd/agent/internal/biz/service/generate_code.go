/*
 *
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package service

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/cloudwego/cwgo/platform/server/cmd/agent/internal/svc"
	"github.com/cloudwego/cwgo/platform/server/shared/consts"
	"github.com/cloudwego/cwgo/platform/server/shared/errx"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/agent"
	"github.com/cloudwego/cwgo/platform/server/shared/logger"
	"github.com/cloudwego/cwgo/platform/server/shared/utils"
	"go.uber.org/zap"
)

type GenerateCodeService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
} // NewGenerateCodeService new GenerateCodeService
func NewGenerateCodeService(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCodeService {
	return &GenerateCodeService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Run create note info
func (s *GenerateCodeService) Run(req *agent.GenerateCodeReq) (resp *agent.GenerateCodeRes, err error) {
	// get idl info by idl id
	idlModel, err := s.svcCtx.DaoManager.Idl.GetIDL(s.ctx, req.IdlId)
	if err != nil {
		logger.Logger.Error("get idl info failed", zap.Error(err))
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumDatabase,
			Msg:  "get idl info failed",
		}, nil
	}

	// get repository info by repository id
	idlRepoModel, err := s.svcCtx.DaoManager.Repository.GetRepository(s.ctx, idlModel.IdlRepositoryId)
	if err != nil {
		logger.Logger.Error("get repo info failed", zap.Error(err))
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumDatabase,
			Msg:  "get repo info failed",
		}, nil
	}

	// get repo client
	repoClient, err := s.svcCtx.RepoManager.GetClient(idlRepoModel.Id)
	if err != nil {
		if errx.GetCode(err) == consts.ErrNumTokenInvalid {
			// repo token is invalid or expired
			return &agent.GenerateCodeRes{
				Code: consts.ErrNumTokenInvalid,
				Msg:  err.Error(),
			}, nil
		}

		logger.Logger.Error(consts.ErrMsgRepoGetClient, zap.Error(err), zap.Int64("repo_id", idlRepoModel.Id))
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumRepoGetClient,
			Msg:  consts.ErrMsgRepoGetClient,
		}, nil
	}

	// parsing URLs to obtain information
	idlPid, owner, repoName, err := repoClient.ParseFileUrl(
		utils.GetRepoFullUrl(
			idlRepoModel.RepositoryType,
			fmt.Sprintf("https://%s/%s/%s",
				idlRepoModel.RepositoryDomain,
				idlRepoModel.RepositoryOwner,
				idlRepoModel.RepositoryName,
			),
			idlRepoModel.RepositoryBranch,
			idlModel.MainIdlPath,
		),
	)
	if err != nil {
		logger.Logger.Error(consts.ErrMsgParamRepositoryUrl, zap.Error(err))
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumParamRepositoryUrl,
			Msg:  consts.ErrMsgParamRepositoryUrl,
		}, nil
	}
	// get the entire repository archive
	archiveData, err := repoClient.GetRepositoryArchive(owner, repoName, idlRepoModel.RepositoryBranch)
	if err != nil {
		logger.Logger.Error(consts.ErrMsgRepoGetArchive, zap.Error(err))
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumRepoGetArchive,
			Msg:  consts.ErrMsgRepoGetArchive,
		}, nil
	}

	// create temp dir
	tempDir, err := os.MkdirTemp(consts.TempDir, strconv.FormatInt(idlRepoModel.Id, 10))
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(consts.TempDir, 0o700)
			if err != nil {
				logger.Logger.Error(consts.ErrMsgCommonCreateTempDir, zap.Error(err))
				return &agent.GenerateCodeRes{
					Code: consts.ErrNumCommonCreateTempDir,
					Msg:  consts.ErrMsgCommonCreateTempDir,
				}, nil
			}

			tempDir, err = os.MkdirTemp(consts.TempDir, strconv.FormatInt(idlRepoModel.Id, 10))
			if err != nil {
				logger.Logger.Error(consts.ErrMsgCommonCreateTempDir, zap.Error(err))
				return &agent.GenerateCodeRes{
					Code: consts.ErrNumCommonCreateTempDir,
					Msg:  consts.ErrMsgCommonCreateTempDir,
				}, nil
			}
		} else {
			logger.Logger.Error(consts.ErrMsgCommonCreateTempDir, zap.Error(err))
			return &agent.GenerateCodeRes{
				Code: consts.ErrNumCommonCreateTempDir,
				Msg:  consts.ErrMsgCommonCreateTempDir,
			}, nil
		}
	}
	defer os.RemoveAll(tempDir)

	tempDirRepo := tempDir + "/" + consts.TempDirRepo

	// the archive type of GitHub is tarball instead of tar
	isTarBall := false
	if idlRepoModel.RepositoryType == consts.RepositoryTypeNumGithub {
		isTarBall = true
	}

	// extract the tar package and persist it to a temporary file
	archiveName, err := utils.UnTar(archiveData, tempDirRepo, isTarBall)
	if err != nil {
		logger.Logger.Error(consts.ErrMsgRepoParseArchive, zap.Error(err))
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumRepoParseArchive,
			Msg:  consts.ErrMsgRepoParseArchive,
		}, nil
	}

	err = os.Mkdir(tempDir+"/"+consts.TempDirGeneratedCode, 0755)
	if err != nil {
		logger.Logger.Error(consts.ErrMsgCommonMkdir, zap.Error(err))
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumCommonMkdir,
			Msg:  consts.ErrMsgCommonMkdir,
		}, nil
	}

	tempDirGeneratedCode := tempDir + "/" + consts.TempDirGeneratedCode

	// generate code using cwgo
	err = s.svcCtx.Generator.Generate(
		idlRepoModel.RepositoryDomain,
		idlRepoModel.RepositoryOwner,
		tempDirRepo+"/"+archiveName+idlPid,
		idlModel.ServiceName,
		tempDirGeneratedCode,
	)
	if err != nil {
		logger.Logger.Error(consts.ErrMsgCommonGenerateCode, zap.Error(err))
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumCommonGenerateCode,
			Msg:  consts.ErrMsgCommonGenerateCode,
		}, nil
	}

	fileContentMap := make(map[string][]byte)
	// parse the file and add it to the map
	if err := utils.ProcessFolders(
		fileContentMap,
		tempDirGeneratedCode,
		"kitex_gen", "rpc", "go.mod", "go.sum",
	); err != nil {
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumCommonProcessFolders,
			Msg:  consts.ErrMsgCommonProcessFolders,
		}, nil
	}

	// push files to the repository
	serviceRepositoryModel, err := s.svcCtx.DaoManager.Repository.GetRepository(s.ctx, idlModel.ServiceRepositoryId)
	if err != nil {
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumDatabase,
			Msg:  consts.ErrMsgDatabase,
		}, nil
	}

	err = repoClient.PushFilesToRepository(fileContentMap, owner, serviceRepositoryModel.RepositoryName, consts.MainRef, "generated by cwgo")
	if err != nil {
		return &agent.GenerateCodeRes{
			Code: consts.ErrNumRepoPush,
			Msg:  consts.ErrMsgRepoPush,
		}, nil
	}

	return &agent.GenerateCodeRes{
		Code: 0,
		Msg:  "generate code successfully",
	}, nil
}
