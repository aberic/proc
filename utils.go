package proc

import (
	"github.com/aberic/gnomon"
	"github.com/aberic/gnomon/log"
	"os"
)

const (
	procDir         = "PROC_DIR"
	ListenAddr      = "LISTEN_ADDR"
	hostname        = "HOSTNAME"
	timeDistanceEnv = "TIME_DISTANCE"
)

// FileRootPath 读取文件根路径
func FileRootPath() string {
	return gnomon.EnvGetD(procDir, "/proc")
}

const (
	// CodeSuccess 成功
	CodeSuccess = iota
	// CodeFile 失败
	CodeFile
)

// Resp 通用返回结构
type Resp struct {
	Code   int
	ErrMsg string
	Data   interface{}
}

// ResponseSuccess 返回成功结构
func ResponseSuccess(model interface{}) *Resp {
	return &Resp{
		Code:   CodeSuccess,
		ErrMsg: "",
		Data:   model,
	}
}

// ResponseFail 返回失败结构
func ResponseFail(err error) *Resp {
	return &Resp{
		Code:   CodeFile,
		ErrMsg: err.Error(),
		Data:   nil,
	}
}

const (
	// ProductionEnv 是否生产环境，在生产环境下控制台不会输出任何日志
	ProductionEnv = "PRODUCTION"
	// LogDirEnv 日志文件目录
	LogDirEnv = "LOG_DIR"
	// LogFileMaxSizeEnv 每个日志文件保存的最大尺寸 单位：M
	LogFileMaxSizeEnv = "LOG_FILE_MAX_SIZE"
	// LogFileMaxAgeEnv 文件最多保存多少天
	LogFileMaxAgeEnv = "LOG_FILE_MAX_AGE"
	// LogUtcEnv CST & UTC 时间
	LogUtcEnv = "LOG_UTC"
	// LogLevelEnv 日志级别(debugLevel/infoLevel/warnLevel/ErrorLevel/panicLevel/fatalLevel)
	LogLevelEnv = "LOG_LEVEL"
)

// InitLog 初始化log日志组件
//
// 全局main入口均可调用执行
func InitLog() {
	log.Fit(
		gnomon.EnvGetD(LogLevelEnv, "Debug"),
		gnomon.EnvGetD(LogDirEnv, os.TempDir()),
		gnomon.EnvGetIntD(LogFileMaxSizeEnv, 1024),
		gnomon.EnvGetIntD(LogFileMaxAgeEnv, 7),
		gnomon.EnvGetBool(LogUtcEnv),
		gnomon.EnvGetBool(ProductionEnv))
}
