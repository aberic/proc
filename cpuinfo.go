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

type CpuInfo struct {
	Processor       string   // 逻辑处理器的id(0)
	VendorID        string   // CPU制造商(GenuineIntel)
	CpuFamily       string   // CPU产品系列代号(6)
	Model           string   // CPU属于其系列中的哪一代号(79)
	ModelName       string   // CPU属于的名字、编号、主频(Intel(R) Xeon(R) CPU E5-26xx v4)
	Stepping        string   // CPU属于制作更新版本(1)
	Microcode       string   // (0x1)
	CpuMHz          string   // CPU的实际使用主频(2394.454)
	CacheSize       string   // CPU二级cache大小(4096 KB)
	PhysicalID      string   // 物理封装的处理器的id，从0开始，说明我的服务器有两个物理CPU(0)
	Siblings        string   // 位于相同物理封装的处理器中的逻辑处理器的数量(1)
	CoreID          string   // 当前物理核在其所处的CPU中的编号，该编号不一定连续(0)
	CpuCores        string   // 该逻辑核所处CPU的物理核数(1)
	ApicID          string   // 用来区分不同逻辑和的编号，每个逻辑和的此编号不同，不一定连续(0)
	InitialApicID   string   // (0)
	Fpu             string   // 是否具有浮点运算单元(yes)
	FpuException    string   // 是否支持浮点计算异常(yes)
	CpuIDLevel      string   // 执行cpuid指令前，eax寄存器中的值，不同cpuid指令会返回不同内容(13)
	WP              string   // 表明当前CPU是否在内核态支持对用户空间的写保护(yes)
	Flags           []string // 当前CPU支持的功能(fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss ht syscall nx lm constant_tsc rep_good nopl eagerfpu pni pclmulqdq ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch bmi1 avx2 bmi2 rdseed adx xsaveopt)
	Bogomips        string   // 测算CPU速度(4788.90)
	ClFlushSize     string   // 每次刷新缓存的大小单位(64)
	CacheAlignment  string   // 缓存地址对齐单位(64)
	AddressSizes    string   // 可访问地址空间为数(40 bits physical, 48 bits virtual)
	PowerManagement string   // 电源管理相关
}

func (c *CpuInfo) FormatCpuInfo(filePath string) {
	data, err := file.ReadFileByLine(filePath)
	if nil != err {
		log.Self.Error("read cpu info error", log.Error(err))
	} else {
		c.Processor = str.Trim(strings.Split(data[0], ":")[1])
		c.VendorID = str.Trim(strings.Split(data[1], ":")[1])
		c.CpuFamily = str.Trim(strings.Split(data[2], ":")[1])
		c.Model = str.Trim(strings.Split(data[3], ":")[1])
		c.ModelName = str.Trim(strings.Split(data[4], ":")[1])
		c.Stepping = str.Trim(strings.Split(data[5], ":")[1])
		c.Microcode = str.Trim(strings.Split(data[6], ":")[1])
		c.CpuMHz = str.Trim(strings.Split(data[7], ":")[1])
		c.CacheSize = str.Trim(strings.Split(data[8], ":")[1])
		c.PhysicalID = str.Trim(strings.Split(data[9], ":")[1])
		c.Siblings = str.Trim(strings.Split(data[10], ":")[1])
		c.CoreID = str.Trim(strings.Split(data[11], ":")[1])
		c.CpuCores = str.Trim(strings.Split(data[12], ":")[1])
		c.ApicID = str.Trim(strings.Split(data[13], ":")[1])
		c.InitialApicID = str.Trim(strings.Split(data[14], ":")[1])
		c.Fpu = str.Trim(strings.Split(data[15], ":")[1])
		c.FpuException = str.Trim(strings.Split(data[16], ":")[1])
		c.CpuIDLevel = str.Trim(strings.Split(data[17], ":")[1])
		c.WP = str.Trim(strings.Split(data[18], ":")[1])
		c.Flags = strings.Split(strings.Split(data[19], ":")[1], " ")
		c.Bogomips = str.Trim(strings.Split(data[20], ":")[1])
		c.ClFlushSize = str.Trim(strings.Split(data[21], ":")[1])
		c.CacheAlignment = str.Trim(strings.Split(data[22], ":")[1])
		c.AddressSizes = str.Trim(strings.Split(data[23], ":")[1])
	}
}
