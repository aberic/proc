package proc

import (
	"fmt"
	"github.com/aberic/gnomon"
	"io/ioutil"
	"time"
)

var (
	proc         *Proc
	remote, host string
	scheduled    *time.Timer // 超时检查对象
	delay        = time.Second * time.Duration(5)
	stop         chan struct{} // 释放当前角色chan
)

func init() {
	proc = &Proc{}
	remote = gnomon.EnvGet(listenAddr)
	host = gnomon.EnvGet(hostname)
	fmt.Println("remote: ", remote)
	scheduled = time.NewTimer(delay)
	stop = make(chan struct{}, 1)
}

// ListenStart 开启监听发送
func ListenStart() {
	if gnomon.StringIsNotEmpty(remote) {
		go send()
	}
}

func send() {
	scheduled.Reset(time.Millisecond * time.Duration(5))
	for {
		select {
		case <-scheduled.C:
			proc.run()
			_, _ = gnomon.HTTPPostJSON(remote, proc)
			scheduled.Reset(delay)
		case <-stop:
			return
		}
	}
}

// Proc 监听发送完整对象
type Proc struct {
	Hostname string
	CPUGroup *CPUGroup
	MemInfo  *MemInfo
	LoadAvg  *LoadAvg
	//Swaps    *Swaps
	Version *Version
	Stat    *Stat
	//CGroup   *CGroup
	UsageCPU float64
}

func (p *Proc) run() {
	cpuGroup := &CPUGroup{CPUArray: []*CPUInfo{}}
	if err := cpuGroup.Info(); nil == err {
		p.CPUGroup = cpuGroup
	}
	memInfo := &MemInfo{}
	if err := memInfo.Info(); nil == err {
		p.MemInfo = memInfo
	}
	loadAvg := &LoadAvg{}
	if err := loadAvg.Info(); nil == err {
		p.LoadAvg = loadAvg
	}
	//swaps := &Swaps{}
	//if err := swaps.Info(); nil == err {
	//	p.Swaps = swaps
	//}
	version := &Version{}
	if err := version.Info(); nil == err {
		p.Version = version
	}
	stat := &Stat{}
	if err := stat.Info(); nil == err {
		p.Stat = stat
	}
	//cGroup := &CGroup{}
	//if err := cGroup.Info(); nil == err {
	//	p.CGroup = cGroup
	//}
	if usage, err := UsageCPU(); nil == err {
		p.UsageCPU = usage
	}
	bs, err := ioutil.ReadFile(host)
	if nil == err {
		p.Hostname = gnomon.StringTrim(string(bs))
	} else {
		p.Hostname = gnomon.EnvGet("HOSTNAME")
	}
}
