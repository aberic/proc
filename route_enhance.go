package proc

import (
	"github.com/aberic/gnomon/grope"
	"github.com/aberic/proc/comm"
	"net/http"
)

// RouterProc 路由
func RouterEnhance(hs *grope.GHttpServe) {
	// 仓库相关路由设置
	route := hs.Group("/enhance")
	route.Get("/cpu/usage", &CPUInfo{}, cpuUsage)
}

func cpuUsage(_ http.ResponseWriter, _ *http.Request, _ interface{}, _ map[string]string) (respModel interface{}, custom bool) {
	var (
		usage float64
		err   error
	)
	if usage, err = UsageCPU(); nil != err {
		return comm.ResponseFail(err), false
	}
	return comm.ResponseSuccess(usage), false
}
