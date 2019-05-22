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
	"github.com/ennoo/rivet"
	"github.com/ennoo/rivet/trans/response"
	"strings"
)

const (
	// FileRootPath 读取文件根路径
	FileRootPath = "/proc"
)

//var FileRootPath = strings.Join([]string{env.GetEnv("GOPATH"), "/src/github.com/ennoo/proc/files"}, "")

// RouterProc 路由
func RouterProc(router *response.Router) {
	// 仓库相关路由设置
	router.Group = router.Engine.Group("/proc")
	router.GET("/cpu", cpu)
	router.GET("/mem", mem)
	router.GET("/loadavg", loadavg)
	router.GET("/swaps", swaps)
	router.GET("/version", version)
	router.GET("/stat", stat)
	router.GET("/cpu/usage", cpuUsage)
}

func cpu(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		cpuInfo := CPUInfo{}
		cpuInfo.FormatCPUInfo(strings.Join([]string{FileRootPath, "/cpuinfo"}, ""))
		result.SaySuccess(router.Context, cpuInfo)
	})
}

func mem(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		memInfo := MemInfo{}
		memInfo.FormatMemInfo(strings.Join([]string{FileRootPath, "/meminfo"}, ""))
		result.SaySuccess(router.Context, memInfo)
	})
}

func loadavg(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		loadAvg := LoadAvg{}
		loadAvg.FormatLoadAvg(strings.Join([]string{FileRootPath, "/loadavg"}, ""))
		result.SaySuccess(router.Context, loadAvg)
	})
}

func swaps(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		swaps := Swaps{}
		swaps.FormatSwaps(strings.Join([]string{FileRootPath, "/swaps"}, ""))
		result.SaySuccess(router.Context, swaps)
	})
}

func version(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		version := Version{}
		version.FormatVersion(strings.Join([]string{FileRootPath, "/version"}, ""))
		result.SaySuccess(router.Context, version)
	})
}

func stat(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		stat := Stat{}
		stat.FormatStat(strings.Join([]string{FileRootPath, "/stat"}, ""))
		result.SaySuccess(router.Context, stat)
	})
}

func cpuUsage(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		result.SaySuccess(router.Context, UsageCPU())
	})
}
