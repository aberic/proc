/*
 * Copyright (c) 2020. aberic - All Rights Reserved.
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
	"bufio"
	"github.com/aberic/gnomon"
	"os"
	"strings"
	"sync"
)

var (
	MountsInstance     *Mounts
	MountsInstanceOnce sync.Once
)

type Mounts struct {
	Mounts []Mount `json:"mounts"`
}

type Mount struct {
	Device     string `json:"device"`
	MountPoint string `json:"mountpoint"`
	FSType     string `json:"fstype"`
	Options    string `json:"options"`
}

func obtainMounts() *Mounts {
	MountsInstanceOnce.Do(func() {
		if nil == MountsInstance {
			MountsInstance = &Mounts{Mounts: []Mount{}}
		}
	})
	return MountsInstance
}

const (
	DefaultBufferSize = 1024
)

// Info DiskStats 对象
func (m *Mounts) Info() error {
	return m.readMounts(gnomon.StringBuild(FileRootPath(), "/mounts"))
}

func (m *Mounts) readMounts(path string) error {
	fin, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = fin.Close() }()

	m.Mounts = []Mount{}

	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		var mount = &Mount{
			fields[0],
			fields[1],
			fields[2],
			fields[3],
		}
		m.Mounts = append(m.Mounts, *mount)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
