package proc

import (
	"github.com/ennoo/rivet"
	"github.com/ennoo/rivet/trans/response"
)

// RouterEnhance 路由
func RouterEnhance(router *response.Router) {
	// 仓库相关路由设置
	router.Group = router.Engine.Group("/enhance")
	router.GET("/cpu/usage", cpuUsage)
}

func cpuUsage(router *response.Router) {
	rivet.Response().Do(router.Context, func(result *response.Result) {
		result.SaySuccess(router.Context, UsageCPU())
	})
}
