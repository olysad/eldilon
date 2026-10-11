package errs

import (
	"errors"
	"strings"
)

// type define

type buildErrStruct struct {
	templ string
	args  []any
}

// func define

// IsErr 收口依赖 对标准Is封装
func IsErr(err error, target error) bool {
	return errors.Is(err, target)
}

func (b *buildErrStruct) builds() error {
	openData := *b
	v := strings.SplitAfter(openData.templ, "$")
	var result []string
	for i := range len(v) {
		result = append(result, goEx(v[i], string(openData.args[i].(string))))
	}
	return errors.New(strings.Join(result, ""))
}

// BuildErr 用 $ 作为占位符构造 error。
//
// 注意：本实现为手写模板引擎，简单场景够用，
// 参数个数与占位符不匹配会 panic。
// 新代码建议直接用 fmt.Errorf("...%v...", args...)。
//
// TODO(me): 有朝一日重构为 fmt.Errorf
// 作者：当年的我，心血之作，勿删。
// NOTE Deprecated: 即将废弃！！！
func BuildErr(templ string, args ...any) error {
	templStruct := &buildErrStruct{
		templ, append(args, ""),
	}
	return templStruct.builds()
}

func goEx(str string, target string) string {
	return strings.Replace(str, "$", target, -1)
}
