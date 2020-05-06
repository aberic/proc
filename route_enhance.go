package proc

import (
	"github.com/aberic/gnomon/grope"
	"net/http"
)

// RouterEnhance 路由
func RouterEnhance(hs *grope.GHttpServe) {
	// 仓库相关路由设置
	route := hs.Group("/enhance")
	route.Get("/cpu/usage", cpuUsage)
}

func cpuUsage(ctx *grope.Context) {
	var (
		usage float64
		err   error
	)
	if usage, err = UsageCPU(); nil != err {
		_ = ctx.ResponseJSON(http.StatusBadRequest, ResponseFail(err))
	}
	_ = ctx.ResponseJSON(http.StatusOK, ResponseSuccess(usage))
}
