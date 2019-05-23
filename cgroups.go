package proc

import (
	"github.com/ennoo/rivet/utils/file"
	"github.com/ennoo/rivet/utils/log"
	"github.com/ennoo/rivet/utils/string"
	"strconv"
	"strings"
)

// CGroups CGroup集合
var CGroups []CGroup

// CGroup 是Linux下的一种将进程按组进行管理的机制，在用户层看来，cgroup技术就是把系统中的所有进程组织成一颗一颗独立的树，
// 每棵树都包含系统的所有进程，树的每个节点是一个进程组，而每颗树又和一个或者多个subsystem关联，树的作用是将进程分组，
// 而subsystem的作用就是对这些组进行操作。
type CGroup struct {
	// SubSysName subsystem的名字
	SubSysName string
	// subsystem所关联到的cgroup树的ID，如果多个subsystem关联到同一颗cgroup树，那么他们的这个字段将一样，
	// 比如这里的cpu和cpuacct就一样，表示他们绑定到了同一颗树。如果出现下面的情况，这个字段将为0：
	//
	// 当前subsystem没有和任何cgroup树绑定
	//
	// 当前subsystem已经和cgroup v2的树绑定
	//
	// 当前subsystem没有被内核开启
	Hierarchy int
	// subsystem所关联的cgroup树中进程组的个数，也即树上节点的个数
	NumCGroups int
	// 1表示开启，0表示没有被开启(可以通过设置内核的启动参数“cgroup_disable”来控制subsystem的开启).
	Enabled bool
}

// FormatCGroups 将文件内容转为 CPUInfo 对象
func (c *CGroup) FormatCGroups(filePath string) {
	data, err := file.ReadFileByLine(filePath)
	if nil != err {
		log.Self.Error("read cpu info error", log.Error(err))
	} else {
		size := len(data)
		CGroups = make([]CGroup, size-1)
		for i := 1; i < size; i++ {
			c.formatCGroups(strings.Split(str.SingleSpace(data[i]), " "))
			CGroups[i-1] = *c
		}
	}
}

func (c *CGroup) formatCGroups(arr []string) {
	c.SubSysName = arr[0]
	c.Hierarchy, _ = strconv.Atoi(arr[1])
	c.NumCGroups, _ = strconv.Atoi(arr[2])
	if enable := arr[3]; str.IsNotEmpty(enable) && enable == "1" {
		c.Enabled = true
	} else {
		c.Enabled = false
	}
}
