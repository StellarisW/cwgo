/*
 *
 *  * Copyright 2022 CloudWeGo Authors
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package idl

import (
	"context"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/biz/model/idl"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/svc"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/agent"
	"github.com/cloudwego/cwgo/platform/server/shared/logger"
	"go.uber.org/zap"
	"net/http"
)

const (
	successMsgDeleteIDL = "delete idls successfully"
)

type DeleteIDLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteIDLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteIDLLogic {
	return &DeleteIDLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteIDLLogic) DeleteIDL(req *idl.DeleteIDLsReq) (res *idl.DeleteIDLsRes) {
	client, err := l.svcCtx.Manager.GetAgentClient()
	if err != nil {
		logger.Logger.Error("get rpc client failed", zap.Error(err))
		return &idl.DeleteIDLsRes{
			Code: http.StatusInternalServerError,
			Msg:  "internal err",
		}
	}

	rpcRes, err := client.DeleteIDL(l.ctx, &agent.DeleteIDLsReq{
		Ids: req.Ids,
	})
	if err != nil {
		logger.Logger.Error("connect to rpc client failed", zap.Error(err))
		return &idl.DeleteIDLsRes{
			Code: http.StatusInternalServerError,
			Msg:  "internal err",
		}
	}
	if rpcRes.Code != 0 {
		return &idl.DeleteIDLsRes{
			Code: http.StatusBadRequest,
			Msg:  rpcRes.Msg,
		}
	}

	return &idl.DeleteIDLsRes{
		Code: 0,
		Msg:  successMsgDeleteIDL,
	}
}
