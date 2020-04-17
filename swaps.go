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
	"strings"
)

// Swaps 显示的是交换分区的使用情况
type Swaps struct {
	Filename string
	Type     string
	Size     string
	Used     string
	Priority string
}

// Info Swaps 对象
func (s *Swaps) Info() error {
	return s.doFormatSwaps(strings.Join([]string{comm.FileRootPath(), "/swaps"}, ""))
}

// FormatSwaps 将文件内容转为 Swaps 对象
func (s *Swaps) doFormatSwaps(filePath string) error {
	data, err := gnomon.File().ReadLines(filePath)
	if nil != err {
		return err
	} else {
		swap := gnomon.String().SingleSpace(data[1])
		swaps := strings.Split(swap, " ")
		s.Filename = swaps[0]
		s.Type = swaps[1]
		s.Size = swaps[2]
		s.Used = swaps[3]
		s.Priority = swaps[4]
	}
	return nil
}
