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

package main

import (
	"github.com/ennoo/rivet/utils/file"
	"github.com/ennoo/rivet/utils/log"
	"github.com/ennoo/rivet/utils/string"
	"strings"
)

// CPUInfo 中央处理器信息
type CPUInfo struct {
	Processor       string   // 逻辑处理器的id(0)
	VendorID        string   // CPU制造商(GenuineIntel)
	CPUFamily       string   // CPU产品系列代号(6)
	Model           string   // CPU属于其系列中的哪一代号(79)
	ModelName       string   // CPU属于的名字、编号、主频(Intel(R) Xeon(R) CPU E5-26xx v4)
	Stepping        string   // CPU属于制作更新版本(1)
	Microcode       string   // (0x1)
	CPUMHz          string   // CPU的实际使用主频(2394.454)
	CacheSize       string   // CPU二级cache大小(4096 KB)
	PhysicalID      string   // 物理封装的处理器的id，从0开始，说明我的服务器有两个物理CPU(0)
	Siblings        string   // 位于相同物理封装的处理器中的逻辑处理器的数量(1)
	CoreID          string   // 当前物理核在其所处的CPU中的编号，该编号不一定连续(0)
	CPUCores        string   // 该逻辑核所处CPU的物理核数(1)
	ApicID          string   // 用来区分不同逻辑和的编号，每个逻辑和的此编号不同，不一定连续(0)
	InitialApicID   string   // (0)
	Fpu             string   // 是否具有浮点运算单元(yes)
	FpuException    string   // 是否支持浮点计算异常(yes)
	CPUIDLevel      string   // 执行cpuid指令前，eax寄存器中的值，不同cpuid指令会返回不同内容(13)
	WP              string   // 表明当前CPU是否在内核态支持对用户空间的写保护(yes)
	Flags           []string // 当前CPU支持的功能(fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss ht syscall nx lm constant_tsc rep_good nopl eagerfpu pni pclmulqdq ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch bmi1 avx2 bmi2 rdseed adx xsaveopt)
	Bogomips        string   // 测算CPU速度(4788.90)
	ClFlushSize     string   // 每次刷新缓存的大小单位(64)
	CacheAlignment  string   // 缓存地址对齐单位(64)
	AddressSizes    string   // 可访问地址空间为数(40 bits physical, 48 bits virtual)
	PowerManagement string   // 电源管理相关
}

// FormatCPUInfo 将文件内容转为 CPUInfo 对象
func (c *CPUInfo) FormatCPUInfo(filePath string) {
	data, err := file.ReadFileByLine(filePath)
	if nil != err {
		log.Self.Error("read cpu info error", log.Error(err))
	} else {
		for index := range data {
			c.formatCPUInfo(data[index])
		}
	}
}

func (c *CPUInfo) formatCPUInfo(lineStr string) {
	if strings.HasPrefix(lineStr, "processor") {
		c.Processor = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "vendor") {
		c.VendorID = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cpu f") {
		c.CPUFamily = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "model n") {
		c.ModelName = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "model") {
		c.Model = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "stepping") {
		c.Stepping = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "microcode") {
		c.Microcode = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cpu M") {
		c.CPUMHz = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cache s") {
		c.CacheSize = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "physical") {
		c.PhysicalID = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "siblings") {
		c.Siblings = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "core") {
		c.CoreID = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cpu c") {
		c.CPUCores = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "apicid") {
		c.ApicID = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "initial") {
		c.InitialApicID = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "fpu_exception") {
		c.FpuException = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "fpu") {
		c.Fpu = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cpuid") {
		c.CPUIDLevel = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "wp") {
		c.WP = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "flags") {
		c.Flags = strings.Split(strings.Split(lineStr, ":")[1], " ")
	} else if strings.HasPrefix(lineStr, "bogomips") {
		c.Bogomips = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "clflush") {
		c.ClFlushSize = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cache_") {
		c.CacheAlignment = str.Trim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "address sizes") {
		c.AddressSizes = str.Trim(strings.Split(lineStr, ":")[1])
	}
}
