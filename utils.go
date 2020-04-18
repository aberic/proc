package proc

// FileRootPath 读取文件根路径
func FileRootPath() string {
	return "/proc"
}

const (
	CodeSuccess = iota
	CodeFile
)

type Resp struct {
	code int
	msg  string
	data interface{}
}

func ResponseSuccess(model interface{}) *Resp {
	return &Resp{
		code: CodeSuccess,
		msg:  "",
		data: model,
	}
}

func ResponseFail(err error) *Resp {
	return &Resp{
		code: CodeFile,
		msg:  err.Error(),
		data: nil,
	}
}
