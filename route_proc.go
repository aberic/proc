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
	"github.com/aberic/proc/comm"
	"net/http"
)

// RouterProc 路由
func RouterProc(hs *grope.GHttpServe) {
	// 仓库相关路由设置
	route := hs.Group("/proc")
	route.Get("/cpu", &CPUInfo{}, cpu)
	route.Get("/mem", &MemInfo{}, mem)
	route.Get("/loadavg", &LoadAvg{}, loadavg)
	route.Get("/swaps", &Swaps{}, swaps)
	route.Get("/version", &Version{}, version)
	route.Get("/stat", &Stat{}, stat)
	route.Get("/cgroups", &CGroup{}, cGroups)
}

func cpu(_ http.ResponseWriter, _ *http.Request, _ interface{}, _ map[string]string) (respModel interface{}, custom bool) {
	cpuInfo := CPUInfo{}
	if err := cpuInfo.Info(); nil != err {
		return comm.ResponseFail(err), false
	}
	return comm.ResponseSuccess(cpuInfo), false
}

func mem(_ http.ResponseWriter, _ *http.Request, _ interface{}, _ map[string]string) (respModel interface{}, custom bool) {
	memInfo := MemInfo{}
	if err := memInfo.Info(); nil != err {
		return comm.ResponseFail(err), false
	}
	return comm.ResponseSuccess(memInfo), false
}

func loadavg(_ http.ResponseWriter, _ *http.Request, _ interface{}, _ map[string]string) (respModel interface{}, custom bool) {
	loadAvg := LoadAvg{}
	if err := loadAvg.Info(); nil != err {
		return comm.ResponseFail(err), false
	}
	return comm.ResponseSuccess(loadAvg), false
}

func swaps(_ http.ResponseWriter, _ *http.Request, _ interface{}, _ map[string]string) (respModel interface{}, custom bool) {
	swaps := Swaps{}
	if err := swaps.Info(); nil != err {
		return comm.ResponseFail(err), false
	}
	return comm.ResponseSuccess(swaps), false
}

func version(_ http.ResponseWriter, _ *http.Request, _ interface{}, _ map[string]string) (respModel interface{}, custom bool) {
	version := Version{}
	if err := version.Info(); nil != err {
		return comm.ResponseFail(err), false
	}
	return comm.ResponseSuccess(version), false
}

func stat(_ http.ResponseWriter, _ *http.Request, _ interface{}, _ map[string]string) (respModel interface{}, custom bool) {
	stat := Stat{}
	if err := stat.Info(); nil != err {
		return comm.ResponseFail(err), false
	}
	return comm.ResponseSuccess(stat), false
}

func cGroups(_ http.ResponseWriter, _ *http.Request, _ interface{}, _ map[string]string) (respModel interface{}, custom bool) {
	cGroup := CGroup{}
	if err := cGroup.Info(); nil != err {
		return comm.ResponseFail(err), false
	}
	return comm.ResponseSuccess(cGroup), false
}
