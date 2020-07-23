package proc

import (
	"context"
	"encoding/json"
	"github.com/aberic/gnomon"
	"github.com/aberic/gnomon/log"
	"github.com/aberic/proc/protos"
	"google.golang.org/grpc"
	"io/ioutil"
	"time"
)

var (
	proc      *Proc
	host      string
	scheduled *time.Timer // 超时检查对象
	delay     time.Duration
	stop      chan struct{} // 释放当前角色chan
)

func init() {
	proc = &Proc{}
	timeDistance := gnomon.EnvGetInt64D(timeDistanceEnv, 1500)
	delay = time.Millisecond * time.Duration(timeDistance)
	host = gnomon.EnvGet(hostname)
	scheduled = time.NewTimer(delay)
	stop = make(chan struct{}, 1)
}

// ListenStart 开启监听发送
func ListenStart(remote string, useHTTP bool) {
	log.Debug("listener start", log.Server("proc"), log.Field("remote", remote))
	if gnomon.StringIsNotEmpty(remote) {
		go send(remote, useHTTP)
	}
}

// remote swarm:20219
func send(remote string, useHTTP bool) {
	scheduled.Reset(time.Millisecond * time.Duration(5))
	for {
		select {
		case <-scheduled.C:
			if err := proc.run(); nil == err {
				log.Debug("send", log.Server("proc"), log.Field("proc", proc))
				if useHTTP {
					if _, err := gnomon.HTTPPostJSON(remote, proc); nil != err {
						log.Error("send", log.Err(err))
					}
				} else {
					var (
						procBytes []byte
						err       error
					)
					if procBytes, err = json.Marshal(proc); nil != err {
						log.Error("send", log.Err(err))
					} else {
						if _, err := gnomon.GRPCRequestSingleConn(remote, func(conn *grpc.ClientConn) (i interface{}, err error) {
							// 创建grpc客户端
							cli := protos.NewProcClient(conn)
							// 客户端向grpc服务端发起请求
							return cli.Info(context.Background(), &protos.Request{Proc: procBytes})
						}); nil != err {
							log.Error("send sync", log.Err(err))
						}
					}
				}
			} else {
				log.Error("send", log.Err(err))
			}
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
	Mounts   *Mounts
	Disk     *Disk
	//DiskStats *DiskStats
	SockStat *SockStat
}

func (p *Proc) run() error {
	if err := obtainCPUGroup().Info(); nil == err {
		p.CPUGroup = obtainCPUGroup()
	}
	if err := obtainMemInfo().Info(); nil == err {
		p.MemInfo = obtainMemInfo()
	}
	if err := obtainLoadAvg().Info(); nil == err {
		p.LoadAvg = obtainLoadAvg()
	}
	//swaps := &Swaps{}
	//if err := swaps.Info(); nil == err {
	//	p.Swaps = swaps
	//}
	if err := obtainVersion().Info(); nil == err {
		p.Version = obtainVersion()
	}
	if err := obtainStat().Info(); nil == err {
		p.Stat = obtainStat()
	}
	//cGroup := &CGroup{}
	//if err := cGroup.Info(); nil == err {
	//	p.CGroup = cGroup
	//}
	if usage, err := UsageCPU(); nil == err {
		p.UsageCPU = usage
	}
	if err := obtainMounts().Info(); nil == err {
		p.Mounts = obtainMounts()
	}
	if err := obtainDisk().Info(); nil != err {
		p.Disk = obtainDisk()
	}
	//if err := obtainDiskStats().Info(); nil != err {
	//	p.DiskStats = obtainDiskStats()
	//}
	if err := obtainSockStat().Info(); nil != err {
		p.SockStat = obtainSockStat()
	}
	bs, err := ioutil.ReadFile(host)
	if nil != err {
		return err
	}
	p.Hostname = gnomon.StringTrim(string(bs))
	return nil
}
