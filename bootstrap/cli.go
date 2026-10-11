package bootstrap

import (
	. "eldilon/errs"
)

type ot int
type R struct {
	Mode ot
	File string
}

const (
	// RO 表示run模式。
	//
	// 该命令用于触发解释器立即编译并执行的功能
	RO ot = iota
	BO
	HO
	VO
)

// ParseArgs 被main调用的命令解析
//
// 此为第一阶段，用户调用命令执行命令
// 下一阶段为os.ReadFile读取用户书写的源码文件
func ParseArgs(args []string) (R, error) {
	var (
		w string
		t ot
	)
	var a = args[1:]
	//三类指令，compile,run,build
	for u, k := range a {
		switch {
		case k == "run" || k == "r" || k == "-r" || k == "--run":
			if len(a) > u+1 {
				w = a[u+1]
				t = RO
				if w == "" {
					return R{}, ArgNotFoundWithRun
				}
				return R{t, w}, nil
			} else {
				return R{}, BuildErr("期望<$>，找到<$>", "文件名", "空缺")
			}
		case k == "help" || k == "h" || k == "-h" || k == "--help":
			return R{HO, ""}, nil
		case k == "version" || k == "v" || k == "-v" || k == "--version":
			return R{VO, ""}, nil
		default:
			return R{}, BuildErr("意外的 $ 命令", k)
		}
	}
	return R{HO, ""}, nil
}
