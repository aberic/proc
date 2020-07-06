/*
 * Copyright (c) 2020. ENNOO - All Rights Reserved.
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
	"sync"
	"syscall"
)

var (
	diskInstance     *Disk
	diskInstanceOnce sync.Once
)

func obtainDisk() *Disk {
	diskInstanceOnce.Do(func() {
		if nil == diskInstance {
			diskInstance = &Disk{}
		}
	})
	return diskInstance
}

type Disk struct {
	All        uint64
	Used       uint64
	Free       uint64
	FreeInodes uint64
}

func (d *Disk) Info() error {
	fs := syscall.Statfs_t{}
	if err := syscall.Statfs("/", &fs); err != nil {
		return err
	}
	d.All = fs.Blocks * uint64(fs.Bsize)
	d.Free = fs.Bfree * uint64(fs.Bsize)
	d.Used = d.All - d.Free
	d.FreeInodes = fs.Ffree
	return nil
}
