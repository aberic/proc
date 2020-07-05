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
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	sockStatInstance     *SockStat
	sockStatInstanceOnce sync.Once
)

type SockStat struct {
	// sockets:
	SocketsUsed uint64 `json:"sockets_used" field:"sockets.used"`

	// TCP:
	TCPInUse     uint64 `json:"tcp_in_use" field:"TCP.inuse"`
	TCPOrphan    uint64 `json:"tcp_orphan" field:"TCP.orphan"`
	TCPTimeWait  uint64 `json:"tcp_time_wait" field:"TCP.tw"`
	TCPAllocated uint64 `json:"tcp_allocated" field:"TCP.alloc"`
	TCPMemory    uint64 `json:"tcp_memory" field:"TCP.mem"`

	//TCP6:
	TCP6InUse uint64 `json:"tcp6_in_use" field:"TCP6.inuse"`

	// UDP:
	UDPInUse  uint64 `json:"udp_in_use" field:"UDP.inuse"`
	UDPMemory uint64 `json:"udp_memory" field:"UDP.mem"`

	// UDP6:
	UDP6InUse uint64 `json:"udp6_in_use" field:"UDP6.inuse"`

	// UDPLITE:
	UDPLITEInUse uint64 `json:"udplite_in_use" field:"UDPLITE.inuse"`

	// UDPLITE6:
	UDPLITE6InUse uint64 `json:"udplite6_in_use" field:"UDPLITE6.inuse"`

	// RAW:
	RAWInUse uint64 `json:"raw_in_use" field:"RAW.inuse"`

	// RAW6:
	RAW6InUse uint64 `json:"raw6_in_use" field:"RAW6.inuse"`

	// FRAG:
	FRAGInUse  uint64 `json:"frag_in_use" field:"FRAG.inuse"`
	FRAGMemory uint64 `json:"frag_memory" field:"FRAG.memory"`

	// FRAG6:
	FRAG6InUse  uint64 `json:"frag6_in_use" field:"FRAG6.inuse"`
	FRAG6Memory uint64 `json:"frag6_memory" field:"FRAG6.memory"`
}

func obtainSockStat() *SockStat {
	sockStatInstanceOnce.Do(func() {
		if nil == sockStatInstance {
			sockStatInstance = &SockStat{}
		}
	})
	return sockStatInstance
}

// Info DiskStats 对象
func (s *SockStat) Info() error {
	return s.readSockStat(gnomon.StringBuild(FileRootPath(), "/net/sockstat"))
}

func (s *SockStat) readSockStat(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	// Maps a meminfo metric to its value (i.e. MemTotal --> 100000)
	statMap := map[string]uint64{}
	for _, line := range lines {
		if strings.Index(line, ":") == -1 {
			continue
		}

		statType := line[0:strings.Index(line, ":")] + "."

		// The fields have this pattern: inuse 27 orphan 1 tw 23 alloc 31 mem 3
		// The stats are grouped into pairs and need to be parsed and placed into the stat map.
		key := ""
		for k, v := range strings.Fields(line[strings.Index(line, ":")+1:]) {
			// Every second field is a value.
			if (k+1)%2 != 0 {
				key = v
				continue
			}
			val, _ := strconv.ParseUint(v, 10, 64)
			statMap[statType+key] = val
		}
	}

	elem := reflect.ValueOf(s).Elem()
	typeOfElem := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		val, ok := statMap[typeOfElem.Field(i).Tag.Get("field")]
		if ok {
			elem.Field(i).SetUint(val)
		}
	}

	return nil
}
