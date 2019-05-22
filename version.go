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
)

// Version 这个文件只有一行内容，说明正在运行的内核版本。可以用标准的编程方法进行分析获得所需的系统信息
type Version struct {
	version string
}

// FormatVersion 将文件内容转为 Version 对象
func (v *Version) FormatVersion(filePath string) {
	data, err := file.ReadFileFirstLine(filePath)
	if nil != err {
		log.Self.Error("read version error", log.Error(err))
	} else {
		v.version = data
	}
}
