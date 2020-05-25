package proc

import "github.com/aberic/gnomon"

const (
	procDir    = "PROC_DIR"
	listenAddr = "LISTEN_ADDR"
	hostname   = "HOSTNAME"
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
