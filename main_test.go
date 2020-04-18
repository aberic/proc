/*
 * Copyright (c) 2019. aberic - All Rights Reserved.
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
	"github.com/aberic/gnomon"
	"strings"
	"testing"
)

//func TestCpuInfo_FormatCpuInfo(t *testing.T) {
//	cpuInfo := CPUInfo{}
//	cpuInfo.doFormatCPUInfo(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/cpuinfo"}, ""))
//	fmt.Println("cpuInfo:", gnomon.String().ToString(cpuInfo))
//}
//
//func TestMemInfo_FormatMemInfo(t *testing.T) {
//	menInfo := MemInfo{}
//	menInfo.doFormatMemInfo(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/meminfo"}, ""))
//	fmt.Println("menInfo:", gnomon.String().ToString(menInfo))
//}
//
//func TestLoadAvg_FormatLoadAvg(t *testing.T) {
//	loadAvg := LoadAvg{}
//	loadAvg.doFormatLoadAvg(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/loadavg"}, ""))
//	fmt.Println("loadAvg:", gnomon.String().ToString(loadAvg))
//}
//
//func TestSwaps_FormatSwaps(t *testing.T) {
//	swaps := Swaps{}
//	swaps.doFormatSwaps(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/swaps"}, ""))
//	fmt.Println("swaps:", gnomon.String().ToString(swaps))
//}
//
//func TestVersion_FormatVersion(t *testing.T) {
//	version := Version{}
//	version.doFormatVersion(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/version"}, ""))
//
//	fmt.Println("version:", gnomon.String().ToString(version))
//}
//
//func TestStat_FormatStat(t *testing.T) {
//	stat := Stat{}
//	stat.doFormatStat(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/stat"}, ""))
//	fmt.Println("stat:", gnomon.String().ToString(stat))
//}
//
//func TestCGroup_FormatCGroups(t *testing.T) {
//	cGroup := CGroup{}
//	cGroup.doFormatCGroups(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/cgroups"}, ""))
//	fmt.Println("stat:", gnomon.String().ToString(CGroups))
//}

func TestCpuInfo_doFormatCpuInfo(t *testing.T) {
	cpuInfo := CPUInfo{}
	t.Log(cpuInfo.doFormatCPUInfo(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/cpuinfo"}, "")))
	fmt.Println("cpuInfo:", gnomon.String().ToString(cpuInfo))
}

func TestMemInfo_doFormatMemInfo(t *testing.T) {
	menInfo := MemInfo{}
	t.Log(menInfo.doFormatMemInfo(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/meminfo"}, "")))
	fmt.Println("menInfo:", gnomon.String().ToString(menInfo))
}

func TestLoadAvg_doFormatLoadAvg(t *testing.T) {
	loadAvg := LoadAvg{}
	t.Log(loadAvg.doFormatLoadAvg(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/loadavg"}, "")))
	fmt.Println("loadAvg:", gnomon.String().ToString(loadAvg))
}

func TestSwaps_doFormatSwaps(t *testing.T) {
	swaps := Swaps{}
	t.Log(swaps.doFormatSwaps(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/swaps"}, "")))
	fmt.Println("swaps:", gnomon.String().ToString(swaps))
}

func TestVersion_doFormatVersion(t *testing.T) {
	version := Version{}
	t.Log(version.doFormatVersion(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/version"}, "")))
	fmt.Println("version:", gnomon.String().ToString(version))
}

func TestStat_doFormatStat(t *testing.T) {
	stat := Stat{}
	t.Log(stat.doFormatStat(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/stat"}, "")))
	fmt.Println("stat:", gnomon.String().ToString(stat))
}

func TestCGroup_doFormatCGroups(t *testing.T) {
	cGroup := CGroup{}
	t.Log(cGroup.doFormatCGroups(strings.Join([]string{gnomon.Env().Get("GOPATH"), "/src/github.com/aberic/proc/files/cgroups"}, "")))
	fmt.Println("stat:", gnomon.String().ToString(CGroups))
}
