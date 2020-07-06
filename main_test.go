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
	"encoding/json"
	"fmt"
	"github.com/aberic/gnomon"
	"gotest.tools/assert"
	"strings"
	"testing"
)

//func TestCpuInfo_FormatCpuInfo(t *testing.T) {
//	cpuInfo := CPUInfo{}
//	cpuInfo.doFormatCPUInfo(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/cpuinfo"}, ""))
//	fmt.Println("cpuInfo:", gnomon.StringToString(cpuInfo))
//}
//
//func TestMemInfo_FormatMemInfo(t *testing.T) {
//	menInfo := MemInfo{}
//	menInfo.doFormatMemInfo(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/meminfo"}, ""))
//	fmt.Println("menInfo:", gnomon.StringToString(menInfo))
//}
//
//func TestLoadAvg_FormatLoadAvg(t *testing.T) {
//	loadAvg := LoadAvg{}
//	loadAvg.doFormatLoadAvg(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/loadavg"}, ""))
//	fmt.Println("loadAvg:", gnomon.StringToString(loadAvg))
//}
//
//func TestSwaps_FormatSwaps(t *testing.T) {
//	swaps := Swaps{}
//	swaps.doFormatSwaps(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/swaps"}, ""))
//	fmt.Println("swaps:", gnomon.StringToString(swaps))
//}
//
//func TestVersion_FormatVersion(t *testing.T) {
//	version := Version{}
//	version.doFormatVersion(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/version"}, ""))
//
//	fmt.Println("version:", gnomon.StringToString(version))
//}
//
//func TestStat_FormatStat(t *testing.T) {
//	stat := Stat{}
//	stat.doFormatStat(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/stat"}, ""))
//	fmt.Println("stat:", gnomon.StringToString(stat))
//}
//
//func TestCGroup_FormatCGroups(t *testing.T) {
//	cGroup := CGroup{}
//	cGroup.doFormatCGroups(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/cgroups"}, ""))
//	fmt.Println("stat:", gnomon.StringToString(CGroups))
//}

func TestCpuInfo_doFormatCpuInfo(t *testing.T) {
	t.Log(obtainCPUGroup().doFormatCPUGroup(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/cpuinfo"}, "")))
	fmt.Println("cpuInfo:", gnomon.ToString(obtainCPUGroup()))
}

func TestMemInfo_doFormatMemInfo(t *testing.T) {
	t.Log(obtainMemInfo().doFormatMemInfo(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/meminfo"}, "")))
	fmt.Println("menInfo:", gnomon.ToString(obtainMemInfo()))
}

func TestLoadAvg_doFormatLoadAvg(t *testing.T) {
	t.Log(obtainLoadAvg().doFormatLoadAvg(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/loadavg"}, "")))
	fmt.Println("loadAvg:", gnomon.ToString(obtainLoadAvg()))
}

func TestSwaps_doFormatSwaps(t *testing.T) {
	t.Log(obtainSwaps().doFormatSwaps(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/swaps"}, "")))
	fmt.Println("swaps:", gnomon.ToString(obtainSwaps()))
}

func TestVersion_doFormatVersion(t *testing.T) {
	t.Log(obtainVersion().doFormatVersion(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/version"}, "")))
	fmt.Println("version:", gnomon.ToString(obtainVersion()))
}

func TestStat_doFormatStat(t *testing.T) {
	t.Log(obtainStat().doFormatStat(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/stat"}, "")))
	fmt.Println("stat:", gnomon.ToString(obtainStat()))
}

func TestCGroup_doFormatCGroups(t *testing.T) {
	cGroup := CGroup{}
	t.Log(cGroup.doFormatCGroups(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/cgroups"}, "")))
	fmt.Println("stat:", gnomon.ToString(CGroups))
}

func TestDisk_Info(t *testing.T) {
	disk := &Disk{}
	err := disk.read("/")
	assert.NilError(t, err)
	data, err := json.Marshal(disk)
	assert.NilError(t, err)
	t.Log(string(data))
}

func TestDiskStats_ReadDiskStats(t *testing.T) {
	dss := DiskStats{}
	err := dss.readDiskStats(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/diskstat"}, ""))
	assert.NilError(t, err)
	for _, d := range dss.diskStatArr {
		data, err := json.Marshal(d)
		assert.NilError(t, err)
		t.Log(string(data))
		t.Log(d.GetIOTicks(), "|", d.GetReadTicks(), "|", d.GetReadBytes(), "|", d.GetTimeInQueue(), "|", d.GetWriteBytes(), "|", d.GetWriteTicks())
	}
}

func TestMounts_Info(t *testing.T) {
	m := Mounts{}
	err := m.readMounts(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/mounts"}, ""))
	assert.NilError(t, err)
	for _, d := range m.Mounts {
		data, err := json.Marshal(d)
		assert.NilError(t, err)
		t.Log(string(data))
	}
}

func TestNetStat_Info(t *testing.T) {
	n := NetStat{}
	err := n.readNetStat(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/netstat"}, ""))
	assert.NilError(t, err)
	data, err := json.Marshal(n)
	assert.NilError(t, err)
	t.Log(string(data))
}

func TestSockStat_Info(t *testing.T) {
	s := SockStat{}
	err := s.readSockStat(strings.Join([]string{gnomon.EnvGet("GOPATH"), "/src/github.com/aberic/proc/files/sockstat"}, ""))
	assert.NilError(t, err)
	data, err := json.Marshal(s)
	assert.NilError(t, err)
	t.Log(string(data))
}
