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
 */

package proc

import (
	"github.com/aberic/gnomon"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	diskStatInstance     *DiskStats
	diskStatInstanceOnce sync.Once
)

// DiskStat is disk statistics to help measure disk activity.
//
// Note:
// * On a very busy or long-lived system values may wrap.
// * No kernel locks are held while modifying these counters. This implies that
//   minor inaccuracies may occur.
//
// More more info see:
// https://www.kernel.org/doc/Documentation/iostats.txt and
// https://www.kernel.org/doc/Documentation/block/stat.txt
type DiskStat struct {
	Major        int    `json:"major"`         // major device number
	Minor        int    `json:"minor"`         // minor device number
	Name         string `json:"name"`          // device name
	ReadIOs      uint64 `json:"read_ios"`      // number of read I/Os processed
	ReadMerges   uint64 `json:"read_merges"`   // number of read I/Os merged with in-queue I/O
	ReadSectors  uint64 `json:"read_sectors"`  // number of 512 byte sectors read
	ReadTicks    uint64 `json:"read_ticks"`    // total wait time for read requests in milliseconds
	WriteIOs     uint64 `json:"write_ios"`     // number of write I/Os processed
	WriteMerges  uint64 `json:"write_merges"`  // number of write I/Os merged with in-queue I/O
	WriteSectors uint64 `json:"write_sectors"` // number of 512 byte sectors written
	WriteTicks   uint64 `json:"write_ticks"`   // total wait time for write requests in milliseconds
	InFlight     uint64 `json:"in_flight"`     // number of I/Os currently in flight
	IOTicks      uint64 `json:"io_ticks"`      // total time this block device has been active in milliseconds
	TimeInQueue  uint64 `json:"time_in_queue"` // total wait time for all requests in milliseconds
}

type DiskStats struct {
	diskStatArr []DiskStat
}

func obtainDiskStats() *DiskStats {
	diskStatInstanceOnce.Do(func() {
		if nil == diskStatInstance {
			diskStatInstance = &DiskStats{diskStatArr: []DiskStat{}}
		}
	})
	return diskStatInstance
}

// Info DiskStats 对象
func (ds *DiskStats) Info() error {
	return ds.readDiskStats(gnomon.StringBuild(FileRootPath(), "/diskstats"))
}

// ReadDiskStats reads and parses the file.
//
// Note:
// * Assumes a well formed file and will panic if it isn't.
func (ds *DiskStats) readDiskStats(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	devices := strings.Split(string(data), "\n")
	devicesLen := len(devices)
	statArrLen := len(ds.diskStatArr)
	if devicesLen > statArrLen {
		lenDis := devicesLen - statArrLen
		ds.diskStatArr = append(ds.diskStatArr, make([]DiskStat, lenDis)...)
	}

	for i := range ds.diskStatArr {
		fields := strings.Fields(devices[i])
		Major, _ := strconv.ParseInt(fields[0], 10, strconv.IntSize)
		ds.diskStatArr[i].Major = int(Major)
		Minor, _ := strconv.ParseInt(fields[1], 10, strconv.IntSize)
		ds.diskStatArr[i].Minor = int(Minor)
		ds.diskStatArr[i].Name = fields[2]
		ds.diskStatArr[i].ReadIOs, _ = strconv.ParseUint(fields[3], 10, 64)
		ds.diskStatArr[i].ReadMerges, _ = strconv.ParseUint(fields[4], 10, 64)
		ds.diskStatArr[i].ReadSectors, _ = strconv.ParseUint(fields[5], 10, 64)
		ds.diskStatArr[i].ReadTicks, _ = strconv.ParseUint(fields[6], 10, 64)
		ds.diskStatArr[i].WriteIOs, _ = strconv.ParseUint(fields[7], 10, 64)
		ds.diskStatArr[i].WriteMerges, _ = strconv.ParseUint(fields[8], 10, 64)
		ds.diskStatArr[i].WriteSectors, _ = strconv.ParseUint(fields[9], 10, 64)
		ds.diskStatArr[i].WriteTicks, _ = strconv.ParseUint(fields[10], 10, 64)
		ds.diskStatArr[i].InFlight, _ = strconv.ParseUint(fields[11], 10, 64)
		ds.diskStatArr[i].IOTicks, _ = strconv.ParseUint(fields[12], 10, 64)
		ds.diskStatArr[i].TimeInQueue, _ = strconv.ParseUint(fields[13], 10, 64)
	}

	return nil
}

// GetReadBytes 返回读取的字节数
func (d *DiskStat) GetReadBytes() int64 {
	return int64(d.ReadSectors) * 512
}

// GetReadTicks 返回等待读取请求的持续时间
func (d *DiskStat) GetReadTicks() time.Duration {
	return time.Duration(d.ReadTicks) * time.Millisecond
}

// GetWriteBytes 返回写入的字节数
func (d *DiskStat) GetWriteBytes() int64 {
	return int64(d.WriteSectors) * 512
}

// GetWriteTicks 返回写请求等待的持续时间
func (d *DiskStat) GetWriteTicks() time.Duration {
	return time.Duration(d.WriteTicks) * time.Millisecond
}

// GetIOTicks 返回磁盘一直处于活动状态的持续时间
func (d *DiskStat) GetIOTicks() time.Duration {
	return time.Duration(d.IOTicks) * time.Millisecond
}

// GetTimeInQueue 返回所有请求等待的持续时间
func (d *DiskStat) GetTimeInQueue() time.Duration {
	return time.Duration(d.TimeInQueue) * time.Millisecond
}
