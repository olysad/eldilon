package main

import (
	b "eldilon/start"
	utils "eldilon/utils"
	"fmt"
	_ "github.com/spf13/pflag"
)

func main() {
	//参数解析
	R, E := b.ParseArgs(utils.Args())
	utils.TestPrint(R, E, utils.WithSucNotPrint(true))
	switch R.Mode {
	case b.VO:
		fmt.Println(utils.Version)
	case b.HO:
		fmt.Print(utils.Help)
	case b.RO:
		r1, e1 := b.ReadFileByName(R.File)
		utils.TestPrint(r1, e1, utils.WithStatus(utils.Brk))
	}
}
