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
	"github.com/aberic/gnomon"
	"github.com/aberic/proc/comm"
	"strconv"
	"strings"
	"time"
)

// Stat 这个文件包含的信息有 CPU 利用率，磁盘，内存页，内存对换，全部中断，接触开关以及赏赐自举时间
type Stat struct {
	CPUs         []*CPU
	Intr         string // 此处较多冗余信息，简化之，这行给出中断的信息，第一个为自系统启动以来，发生的所有的中断的次数。然后每个数对应一个特定的中断自系统启动以来所发生的次数
	Ctxt         string // 自系统启动以来CPU发生的上下文交换的次数
	BTime        string // 系统启动到现在的时间，单位为秒(s)
	Processes    string // 自系统启动以来所创建的任务的个数目
	ProcsRunning string // 当前运行队列的任务的数目
	ProcsBlocked string // 当前被阻塞的任务的数目
	SoftIrq      string // 此行显示所有CPU的softirq总数，第一列是所有软件和每个软件的总数，后面的列是特定softirq的总数
}

// CPU 利用率
//
// 总的cpu时间totalCpuTime = user + nice + system + idle + iowait + irq + softirq + stealstolen  +  guest
//
// 进程的总Cpu时间processCpuTime = utime + stime + cutime + cstime，该值包括其所有线程的cpu时间
type CPU struct {
	Core      string // CPU核
	User      int64  // 从系统启动开始累计到当前时刻，处于用户态的运行时间，不包含 nice值为负进程
	Nice      int64  // 从系统启动开始累计到当前时刻，nice值为负的进程所占用的CPU时间
	System    int64  // 从系统启动开始累计到当前时刻，处于核心态的运行时间
	Idle      int64  // 从系统启动开始累计到当前时刻，除IO等待时间以外的其它等待时间
	IOWait    int64  // 从系统启动开始累计到当前时刻，IO等待时间(since 2.5.41)
	Irq       int64  // 从系统启动开始累计到当前时刻，硬中断时间(since 2.6.0-test4)
	SoftIrq   int64  // 从系统启动开始累计到当前时刻，软中断时间(since 2.6.0-test4)
	Steal     int64  // 虚拟化环境中运行其他操作系统上花费的时间（since 2.6.11）
	Guest     int64  // 操作系统运行虚拟CPU花费的时间（since 2.6.24）
	GuestNice int64  // 运行一个带nice值的guest花费的时间（since 2.6.33）
}

// Info Stat 对象
func (s *Stat) Info() error {
	return s.doFormatStat(strings.Join([]string{comm.FileRootPath(), "/stat"}, ""))
}

// FormatStat 将文件内容转为 Stat 对象
func (s *Stat) doFormatStat(filePath string) error {
	data, err := gnomon.File().ReadLines(filePath)
	if nil != err {
		return err
	} else {
		for index := range data {
			s.formatStat(data[index])
		}
	}
	return nil
}

func (s *Stat) formatStat(lineStr string) {
	if strings.HasPrefix(lineStr, "cpu") {
		cpuStr := gnomon.String().SingleSpace(lineStr)
		cpuStrArr := strings.Split(cpuStr, " ")
		cpu := CPU{}
		cpu.formatCPU(cpuStrArr)
		s.CPUs = append(s.CPUs, &cpu)
	} else if strings.HasPrefix(lineStr, "intr") {
		s.Intr = lineStr
	} else if strings.HasPrefix(lineStr, "ctxt") {
		s.Ctxt = strings.Split(lineStr, " ")[1]
	} else if strings.HasPrefix(lineStr, "btime") {
		s.BTime = strings.Split(lineStr, " ")[1]
	} else if strings.HasPrefix(lineStr, "processes") {
		s.Processes = strings.Split(lineStr, " ")[1]
	} else if strings.HasPrefix(lineStr, "procs_running") {
		s.ProcsRunning = strings.Split(lineStr, " ")[1]
	} else if strings.HasPrefix(lineStr, "procs_blocked") {
		s.ProcsBlocked = strings.Split(lineStr, " ")[1]
	} else if strings.HasPrefix(lineStr, "softirq") {
		s.SoftIrq = lineStr
	}
}

func (c *CPU) formatCPU(arr []string) {
	c.Core = arr[0]
	var i64 int64
	var err error
	if i64, err = strconv.ParseInt(arr[1], 10, 64); err == nil {
		c.User = i64
	}
	if i64, err = strconv.ParseInt(arr[2], 10, 64); err == nil {
		c.Nice = i64
	}
	if i64, err = strconv.ParseInt(arr[3], 10, 64); err == nil {
		c.System = i64
	}
	if i64, err = strconv.ParseInt(arr[4], 10, 64); err == nil {
		c.Idle = i64
	}
	if i64, err = strconv.ParseInt(arr[5], 10, 64); err == nil {
		c.IOWait = i64
	}
	if i64, err = strconv.ParseInt(arr[6], 10, 64); err == nil {
		c.Irq = i64
	}
	if i64, err = strconv.ParseInt(arr[7], 10, 64); err == nil {
		c.SoftIrq = i64
	}
	if i64, err = strconv.ParseInt(arr[8], 10, 64); err == nil {
		c.Steal = i64
	}
	if i64, err = strconv.ParseInt(arr[9], 10, 64); err == nil {
		c.Guest = i64
	}
	if i64, err = strconv.ParseInt(arr[10], 10, 64); err == nil {
		c.GuestNice = i64
	}
}

// UsageCPU CPU使用率
func UsageCPU() (float64, error) {
	stat1 := Stat{}
	stat2 := Stat{}
	if err := stat1.Info(); nil != err {
		return 0, err
	}
	time.Sleep(10 * time.Millisecond)
	if err := stat2.Info(); nil != err {
		return 0, err
	}
	return usageCPU(stat1.CPUs, stat2.CPUs), nil
}

// usageCPU CPU使用率
func usageCPU(c1s []*CPU, c2s []*CPU) float64 {
	size := int64(len(c1s))
	pcpuTotal := int64(0)
	for i := int64(0); i < size; i++ {
		// 采样两个足够短的时间间隔的Cpu快照，分别计算总的Cpu时间片totalCpuTime
		totalCPUTime1 := c1s[i].User + c1s[i].Nice + c1s[i].System + c1s[i].Idle + c1s[i].IOWait + c1s[i].Irq + c1s[i].SoftIrq + c1s[i].Steal + c1s[i].Guest
		totalCPUTime2 := c2s[i].User + c2s[i].Nice + c2s[i].System + c2s[i].Idle + c2s[i].IOWait + c2s[i].Irq + c2s[i].SoftIrq + c2s[i].Steal + c2s[i].Guest
		// 得到这个时间间隔内的所有时间片
		totalCPUTime := totalCPUTime2 - totalCPUTime1
		// 计算空闲时间idle
		idle := c2s[i].Idle - c1s[i].Idle
		// 计算cpu使用率
		if totalCPUTime <= 0 {
			return 0
		}
		pcpu := 100 * (totalCPUTime - idle) / totalCPUTime
		pcpuTotal += pcpu
	}
	pcpuTotalF64 := gnomon.Scale().Int64toFloat64(pcpuTotal, 2)
	sizeF64 := gnomon.Scale().Int64toFloat64(size, 2)
	return pcpuTotalF64 / sizeF64
}
