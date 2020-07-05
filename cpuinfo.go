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
	"github.com/aberic/gnomon"
	"strings"
	"sync"
)

var (
	cpuGroupInstance     *CPUGroup
	cpuGroupInstanceOnce sync.Once
)

// CPUGroup 中央处理器信息组
type CPUGroup struct {
	CPUArray []*CPUInfo
}

func obtainCPUGroup() *CPUGroup {
	cpuGroupInstanceOnce.Do(func() {
		if nil == cpuGroupInstance {
			cpuGroupInstance = &CPUGroup{CPUArray: []*CPUInfo{}}
		}
	})
	return cpuGroupInstance
}

// Info CPUGroup 对象
func (c *CPUGroup) Info() error {
	return c.doFormatCPUGroup(gnomon.StringBuild(FileRootPath(), "/cpuinfo"))
}

// FormatCPUGroup 将文件内容转为 CPUGroup 对象
func (c *CPUGroup) doFormatCPUGroup(filePath string) error {
	data, err := gnomon.FileReadLines(filePath)
	if nil == err {
		index := 0
		for _, d := range data {
			if gnomon.StringIsEmpty(d) {
				if len(c.CPUArray) <= index || gnomon.StringIsEmpty(c.CPUArray[index].Processor) {
					continue
				}
				index++
				continue
			}
			if len(c.CPUArray) <= index {
				c.CPUArray = append(c.CPUArray, &CPUInfo{})
			}
			c.CPUArray[index].formatCPUInfo(d)
		}
	} else {
		return err
	}
	return nil
}

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

func (c *CPUInfo) formatCPUInfo(lineStr string) {
	if strings.HasPrefix(lineStr, "processor") {
		c.Processor = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "vendor") {
		c.VendorID = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cpu f") {
		c.CPUFamily = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "model n") {
		c.ModelName = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "model") {
		c.Model = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "stepping") {
		c.Stepping = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "microcode") {
		c.Microcode = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cpu M") {
		c.CPUMHz = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cache s") {
		c.CacheSize = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "physical") {
		c.PhysicalID = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "siblings") {
		c.Siblings = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "core") {
		c.CoreID = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cpu c") {
		c.CPUCores = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "apicid") {
		c.ApicID = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "initial") {
		c.InitialApicID = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "fpu_exception") {
		c.FpuException = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "fpu") {
		c.Fpu = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cpuid") {
		c.CPUIDLevel = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "wp") {
		c.WP = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "flags") {
		c.Flags = strings.Split(strings.Split(lineStr, ":")[1], " ")
	} else if strings.HasPrefix(lineStr, "bogomips") {
		c.Bogomips = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "clflush") {
		c.ClFlushSize = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "cache_") {
		c.CacheAlignment = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	} else if strings.HasPrefix(lineStr, "address sizes") {
		c.AddressSizes = gnomon.StringTrim(strings.Split(lineStr, ":")[1])
	}
}
