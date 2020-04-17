package log

import (
	"github.com/aberic/gnomon"
	"os"
	"strings"
)

var (
	logFileDir     = os.TempDir() // 日志文件目录
	logFileMaxSize = 1024         // 每个日志文件保存的最大尺寸 单位：M
	logFileMaxAge  = 7            // 文件最多保存多少天
	//logUtc         = false   // CST & UTC 时间
	logLevel = "Debug" // 日志级别(debugLevel/infoLevel/warnLevel/ErrorLevel/panicLevel/fatalLevel)
	//logProduction  = true   // 是否生产环境，在生产环境下控制台不会输出任何日志
)

func init() {
	if err := initLog(); nil != err {
		panic(err)
	}
}

func initLog() error {
	if err := gnomon.Log().Init(logFileDir, logFileMaxSize, logFileMaxAge, false); nil != err {
		return err
	}
	var level gnomon.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		level = gnomon.Log().DebugLevel()
	case "info":
		level = gnomon.Log().InfoLevel()
	case "warn":
		level = gnomon.Log().WarnLevel()
	case "error":
		level = gnomon.Log().ErrorLevel()
	case "panic":
		level = gnomon.Log().PanicLevel()
	case "fatal":
		level = gnomon.Log().FatalLevel()
	default:
		level = gnomon.Log().DebugLevel()
	}
	gnomon.Log().Set(level, true)
	return nil
}

// Param 日志输出子集对象
type Param struct {
	key   string
	value interface{}
}

func (p *Param) GetKey() string {
	return p.key
}

func (p *Param) GetValue() interface{} {
	return p.value
}

// Field 自定义输出KV对象
func Field(key string, value interface{}) *Param {
	return &Param{key: key, value: value}
}

// Err 自定义输出错误
func Err(err error) *Param {
	if nil != err {
		return &Param{key: "error", value: err.Error()}
	}
	return &Param{key: "error", value: nil}
}

// Errs 自定义输出错误
func Errs(msg string) *Param {
	return &Param{key: "error", value: msg}
}

func Debug(msg string, fields ...gnomon.FieldInter) {
	gnomon.Log().DebugSkip(2, msg, fields...)
}

func Info(msg string, fields ...gnomon.FieldInter) {
	gnomon.Log().InfoSkip(2, msg, fields...)
}

func Warn(msg string, fields ...gnomon.FieldInter) {
	gnomon.Log().WarnSkip(2, msg, fields...)
}

func Error(msg string, fields ...gnomon.FieldInter) {
	gnomon.Log().ErrorSkip(2, msg, fields...)
}

func Panic(msg string, fields ...gnomon.FieldInter) {
	gnomon.Log().PanicSkip(2, msg, fields...)
}

func Fatal(msg string, fields ...gnomon.FieldInter) {
	gnomon.Log().FatalSkip(2, msg, fields...)
}
