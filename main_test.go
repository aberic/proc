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
 *
 */

package proc

import (
	"fmt"
	"github.com/ennoo/rivet/utils/env"
	"github.com/ennoo/rivet/utils/string"
	"strings"
	"testing"
)

func TestCpuInfo_FormatCpuInfo(t *testing.T) {
	cpuInfo := CPUInfo{}
	cpuInfo.FormatCPUInfo(strings.Join([]string{env.GetEnv("GOPATH"), "/src/github.com/ennoo/proc/files/cpuinfo"}, ""))
	fmt.Println("cpuInfo:", str.ToString(cpuInfo))
}

func TestMemInfo_FormatMemInfo(t *testing.T) {
	menInfo := MemInfo{}
	menInfo.FormatMemInfo(strings.Join([]string{env.GetEnv("GOPATH"), "/src/github.com/ennoo/proc/files/meminfo"}, ""))
	fmt.Println("menInfo:", str.ToString(menInfo))
}

func TestLoadAvg_FormatLoadAvg(t *testing.T) {
	loadAvg := LoadAvg{}
	loadAvg.FormatLoadAvg(strings.Join([]string{env.GetEnv("GOPATH"), "/src/github.com/ennoo/proc/files/loadavg"}, ""))
	fmt.Println("loadAvg:", str.ToString(loadAvg))
}

func TestSwaps_FormatSwaps(t *testing.T) {
	swaps := Swaps{}
	swaps.FormatSwaps(strings.Join([]string{env.GetEnv("GOPATH"), "/src/github.com/ennoo/proc/files/swaps"}, ""))
	fmt.Println("swaps:", str.ToString(swaps))
}

func TestVersion_FormatVersion(t *testing.T) {
	version := Version{}
	version.FormatVersion(strings.Join([]string{env.GetEnv("GOPATH"), "/src/github.com/ennoo/proc/files/version"}, ""))

	fmt.Println("version:", str.ToString(version))
}

func TestStat_FormatStat(t *testing.T) {
	stat := Stat{}
	stat.FormatStat(strings.Join([]string{env.GetEnv("GOPATH"), "/src/github.com/ennoo/proc/files/stat"}, ""))
	fmt.Println("stat:", str.ToString(stat))
}

func TestCGroup_FormatCGroups(t *testing.T) {
	cGroup := CGroup{}
	cGroup.FormatCGroups(strings.Join([]string{env.GetEnv("GOPATH"), "/src/github.com/ennoo/proc/files/cgroups"}, ""))
	fmt.Println("stat:", str.ToString(CGroups))
}
