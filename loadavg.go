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
	"github.com/aberic/gnomon"
	"strings"
)

// LoadAvg 系统平均负载均衡
type LoadAvg struct {
	LAvg1     string // 1-分钟平均负载
	LAvg5     string // 5-分钟平均负载
	LAvg15    string // 15-分钟平均负载
	NrRunning string // 分子是正在运行的进程数，分母是进程总数
	LastPid   string // 最大的pid值，包括轻量级进程，即线程
}

// Info LoadAvg 对象
func (l *LoadAvg) Info() error {
	return l.doFormatLoadAvg(gnomon.StringBuild(FileRootPath(), "/loadavg"))
}

// FormatLoadAvg 将文件内容转为 LoadAvg 对象
func (l *LoadAvg) doFormatLoadAvg(filePath string) error {
	data, err := gnomon.FileReadFirstLine(filePath)
	if nil == err {
		ds := strings.Split(data, " ")
		l.LAvg1 = ds[0]
		l.LAvg5 = ds[1]
		l.LAvg15 = ds[2]
		l.NrRunning = ds[3]
		l.LastPid = ds[4]
	} else {
		return err
	}
	return nil
}
