/*
 * Copyright (c) 2019. ENNOO - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proc

import (
	"github.com/aberic/gnomon/grope"
	"github.com/aberic/gnomon/log"
	"net/http"
)

// RouterProc 路由
func RouterProc(hs *grope.GHttpServe) {
	// 仓库相关路由设置
	route := hs.Group("/proc")
	route.Get("/cpu", cpu)
	route.Get("/mem", mem)
	route.Get("/loadavg", loadavg)
	route.Get("/swaps", swaps)
	route.Get("/version", version)
	route.Get("/stat", stat)
	route.Get("/cgroups", cGroups)
	route.Get("/disk", disks)
}

func cpu(ctx *grope.Context) {
	if err := obtainCPUGroup().Info(); nil != err {
		_ = ctx.ResponseJSON(http.StatusBadRequest, ResponseFail(err))
	}
	log.Debug("RouterProc", log.Server("proc"), log.Field("cpu", obtainCPUGroup()))
	_ = ctx.ResponseJSON(http.StatusOK, ResponseSuccess(obtainCPUGroup()))
}

func mem(ctx *grope.Context) {
	if err := obtainMemInfo().Info(); nil != err {
		_ = ctx.ResponseJSON(http.StatusBadRequest, ResponseFail(err))
	}
	log.Debug("RouterProc", log.Server("proc"), log.Field("mem", obtainMemInfo()))
	_ = ctx.ResponseJSON(http.StatusOK, ResponseSuccess(obtainMemInfo()))
}

func loadavg(ctx *grope.Context) {
	if err := obtainLoadAvg().Info(); nil != err {
		_ = ctx.ResponseJSON(http.StatusBadRequest, ResponseFail(err))
	}
	log.Debug("RouterProc", log.Server("proc"), log.Field("loadavg", obtainLoadAvg()))
	_ = ctx.ResponseJSON(http.StatusOK, ResponseSuccess(obtainLoadAvg()))
}

func swaps(ctx *grope.Context) {
	if err := obtainSwaps().Info(); nil != err {
		_ = ctx.ResponseJSON(http.StatusBadRequest, ResponseFail(err))
	}
	log.Debug("RouterProc", log.Server("proc"), log.Field("swaps", obtainSwaps()))
	_ = ctx.ResponseJSON(http.StatusOK, ResponseSuccess(obtainSwaps()))
}

func version(ctx *grope.Context) {
	if err := obtainVersion().Info(); nil != err {
		_ = ctx.ResponseJSON(http.StatusBadRequest, ResponseFail(err))
	}
	log.Debug("RouterProc", log.Server("proc"), log.Field("version", obtainVersion()))
	_ = ctx.ResponseJSON(http.StatusOK, ResponseSuccess(obtainVersion()))
}

func stat(ctx *grope.Context) {
	if err := obtainStat().Info(); nil != err {
		_ = ctx.ResponseJSON(http.StatusBadRequest, ResponseFail(err))
	}
	log.Debug("RouterProc", log.Server("proc"), log.Field("stat", obtainStat()))
	_ = ctx.ResponseJSON(http.StatusOK, ResponseSuccess(obtainStat()))
}

func cGroups(ctx *grope.Context) {
	cGroup := &CGroup{}
	if err := cGroup.Info(); nil != err {
		_ = ctx.ResponseJSON(http.StatusBadRequest, ResponseFail(err))
	}
	log.Debug("RouterProc", log.Server("proc"), log.Field("cGroup", cGroup))
	_ = ctx.ResponseJSON(http.StatusOK, ResponseSuccess(cGroup))
}

func disks(ctx *grope.Context) {
	if err := obtainDisk().Info(); nil != err {
		_ = ctx.ResponseJSON(http.StatusBadRequest, ResponseFail(err))
	}
	log.Debug("RouterProc", log.Server("proc"), log.Field("disk", obtainDisk()))
	_ = ctx.ResponseJSON(http.StatusOK, ResponseSuccess(obtainDisk()))
}
