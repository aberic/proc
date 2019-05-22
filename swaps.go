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

package main

import (
	"github.com/ennoo/rivet/utils/file"
	"github.com/ennoo/rivet/utils/log"
	str "github.com/ennoo/rivet/utils/string"
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

// FormatSwaps 将文件内容转为 Swaps 对象
func (s *Swaps) FormatSwaps(filePath string) {
	data, err := file.ReadFileByLine(filePath)
	if nil != err {
		log.Self.Error("read swaps error", log.Error(err))
	} else {
		swap := str.SingleSpace(data[1])
		swaps := strings.Split(swap, " ")
		s.Filename = swaps[0]
		s.Type = swaps[1]
		s.Size = swaps[2]
		s.Used = swaps[3]
		s.Priority = swaps[4]
	}
}
