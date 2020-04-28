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
}

func cpu(ctx *grope.Context) {
	cpuGroup := &CPUGroup{CPUArray: []*CPUInfo{}}
	if err := cpuGroup.Info(); nil != err {
		_ = ctx.ResponseText(http.StatusBadRequest, err.Error())
	}
	_ = ctx.ResponseJson(http.StatusOK, cpuGroup)
}

func mem(ctx *grope.Context) {
	memInfo := &MemInfo{}
	if err := memInfo.Info(); nil != err {
		_ = ctx.ResponseText(http.StatusBadRequest, err.Error())
	}
	_ = ctx.ResponseJson(http.StatusOK, memInfo)
}

func loadavg(ctx *grope.Context) {
	loadAvg := &LoadAvg{}
	if err := loadAvg.Info(); nil != err {
		_ = ctx.ResponseText(http.StatusBadRequest, err.Error())
	}
	_ = ctx.ResponseJson(http.StatusOK, loadAvg)
}

func swaps(ctx *grope.Context) {
	swaps := &Swaps{}
	if err := swaps.Info(); nil != err {
		_ = ctx.ResponseText(http.StatusBadRequest, err.Error())
	}
	_ = ctx.ResponseJson(http.StatusOK, swaps)
}

func version(ctx *grope.Context) {
	version := &Version{}
	if err := version.Info(); nil != err {
		_ = ctx.ResponseText(http.StatusBadRequest, err.Error())
	}
	_ = ctx.ResponseJson(http.StatusOK, version)
}

func stat(ctx *grope.Context) {
	stat := &Stat{}
	if err := stat.Info(); nil != err {
		_ = ctx.ResponseText(http.StatusBadRequest, err.Error())
	}
	_ = ctx.ResponseJson(http.StatusOK, stat)
}

func cGroups(ctx *grope.Context) {
	cGroup := &CGroup{}
	if err := cGroup.Info(); nil != err {
		_ = ctx.ResponseText(http.StatusBadRequest, err.Error())
	}
	_ = ctx.ResponseJson(http.StatusOK, cGroup)
}
