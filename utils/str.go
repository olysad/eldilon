package utils

var m = ItpMap{
	"version": "0.0.1",
	"run":     "run/r/-r/--run",
	"help":    "help/h/-h/--help",
	"vers":    "version/v/-v/--version",
}

var (
	helpTxt = `欢迎使用 Eldilon 语言解释器！
Eldilon是一门面向句柄开发的语言，其独 特的能力语义、视图检查、句柄机制和漏斗模型使其适用于需要保证访问安全、业务安全和内存安全的中大型项目构建，
当前版本({version})支持的命令如下：
- {run}:
    eldilon {run} <文件名> - 执行一个 Eldilon 源代码文件
- {help}:
    eldilon {help} - 输出帮助信息
- {vers}:
    eldilon {vers} - 打印版本号
`
	Help, _    = m.BuildStrWithInterpol(helpTxt)
	Version, _ = m.BuildStrWithInterpol("版本 {version}")
)
