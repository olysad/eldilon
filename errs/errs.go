package errs

import err "errors"

var (
	ArgNotFoundWithRun   = err.New("期望<文件名>，找到<空缺>")
	ArgNotFound          = err.New("期望<命令>，找到<空缺>")
	ArgNotDefine         = err.New("无法识别并执行非有效指令")
	FileNameIsDir        = err.New("试图非预期的阅读一个目录")
	FileSuffixUnexpected = err.New("不期望或无法识别的文件后缀")
)
