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

package registry

import (
	"context"

	registrylogic "github.com/cloudwego/cwgo/platform/server/cmd/api/internal/biz/logic/registry"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/svc"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/pkg/model/response"
	registrymodel "github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/registry"
	"github.com/cloudwego/cwgo/platform/server/shared/log"
	"github.com/cloudwego/hertz/pkg/app"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.uber.org/zap"
)

// Update .
// @router /registry/update [GET]
func Update(ctx context.Context, rCtx *app.RequestContext) {
	var err error
	var req registrymodel.UpdateReq
	err = rCtx.BindAndValidate(&req)
	if err != nil {
		log.Debug("parse http request failed", zap.Error(err), zap.Reflect("http request", req))
		response.Fail(
			rCtx,
			hertzconsts.StatusBadRequest,
			hertzconsts.StatusBadRequest,
			err.Error(),
		)
		return
	}

	log.Debug("http request args", zap.Reflect("args", req))

	l := registrylogic.NewUpdateLogic(ctx, svc.Svc)

	res := l.Update(&req)

	log.Debug("http response args", zap.Reflect("args", res))

	if res.Code != 0 {
		response.Fail(rCtx, hertzconsts.StatusBadRequest, res.Code, res.Msg)
		return
	}

	response.Ok(rCtx, res.Msg)
}
