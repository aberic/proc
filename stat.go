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
	"github.com/ennoo/rivet/utils/file"
	"github.com/ennoo/rivet/utils/log"
	str "github.com/ennoo/rivet/utils/string"
	"strings"
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
	User      string // 从系统启动开始累计到当前时刻，处于用户态的运行时间，不包含 nice值为负进程
	Nice      string // 从系统启动开始累计到当前时刻，nice值为负的进程所占用的CPU时间
	System    string // 从系统启动开始累计到当前时刻，处于核心态的运行时间
	Idle      string // 从系统启动开始累计到当前时刻，除IO等待时间以外的其它等待时间
	IOWait    string // 从系统启动开始累计到当前时刻，IO等待时间(since 2.5.41)
	Irq       string // 从系统启动开始累计到当前时刻，硬中断时间(since 2.6.0-test4)
	SoftIrq   string // 从系统启动开始累计到当前时刻，软中断时间(since 2.6.0-test4)
	Steal     string // 虚拟化环境中运行其他操作系统上花费的时间（since 2.6.11）
	Guest     string // 操作系统运行虚拟CPU花费的时间（since 2.6.24）
	GuestNice string // 运行一个带nice值的guest花费的时间（since 2.6.33）
}

// FormatStat 将文件内容转为 Stat 对象
func (s *Stat) FormatStat(filePath string) {
	data, err := file.ReadFileByLine(filePath)
	if nil != err {
		log.Self.Error("read stat error", log.Error(err))
	} else {
		for index := range data {
			s.formatStat(data[index])
		}
	}
}

func (s *Stat) formatStat(lineStr string) {
	if strings.HasPrefix(lineStr, "cpu") {
		cpuStr := str.SingleSpace(lineStr)
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
	c.User = arr[1]
	c.Nice = arr[2]
	c.System = arr[3]
	c.Idle = arr[4]
	c.IOWait = arr[5]
	c.Irq = arr[6]
	c.SoftIrq = arr[7]
	c.Steal = arr[8]
	c.Guest = arr[9]
	c.GuestNice = arr[10]
}
