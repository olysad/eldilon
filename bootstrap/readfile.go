package bootstrap

import (
	. "eldilon/errs"
	"os"
	"path/filepath"
	"strings"
)

func ReadFileByName(name string) (any, error) {
	fileStat, err := os.Stat(name)
	if err != nil {
		if IsErr(err, os.ErrNotExist) {
			return "", BuildErr("文件 $ 不存在", name)
		}
	}
	if fileStat.IsDir() {
		return "", BuildErr("输入路径 $ 是一个目录", name)
	} else {
		if strings.EqualFold(filepath.Ext(name), ".eldc") {
			return "", BuildErr("后缀 $ 是一个字节码文件，当前版本暂不支持从字节码解析执行", ".eldc")
		} else if strings.EqualFold(filepath.Ext(name), ".ecp") {
			return "", BuildErr("后缀 .ecp 是一个可执行字节码包，暂不支持执行该文件")
		} else if !strings.EqualFold(filepath.Ext(name), ".eld") {
			return "", BuildErr("输入文件后缀 $ 不是 Eldilon 源代码文件，应为“.eld”后缀文件", filepath.Ext(name))
		}
	}
	R, E := reader(name)
	return string(R), E
}

func reader(name string) ([]byte, error) {
	return os.ReadFile(name)
}
