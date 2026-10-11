package utils

import (
	"fmt"
	"github.com/muesli/termenv"
	"os"
)

type printConfig struct {
	ifSucNotPrint  bool
	ifFailNotPrint bool
	notStyled      bool
	status         programStatus
	clearFunc      func() programStatus
}

type programStatus int

const (
	Ctn programStatus = iota
	Brk
	Ftl
)

type option func(*printConfig)

func WithSucNotPrint(c bool) option {
	return func(arg *printConfig) {
		arg.ifSucNotPrint = c
	}
}

func WithFailNotPrint(c bool) option {
	return func(arg *printConfig) {
		arg.ifFailNotPrint = c
	}
}

func WithNotStyled(c bool) option {
	return func(arg *printConfig) {
		arg.notStyled = c
	}
}

func WithStatus(c programStatus) option {
	return func(arg *printConfig) {
		arg.status = c
	}
}

func WithClearFunc(c func() programStatus) option {
	return func(arg *printConfig) {
		arg.clearFunc = c
	}
}

func TestPrint(R any, E error, ops ...option) {
	zero := &printConfig{}
	output := termenv.NewOutput(os.Stdout)
	for _, v := range ops {
		v(zero)
	}

	if E != nil {
		if !zero.ifFailNotPrint {
			str := fmt.Sprintf("Error:%v\n", E)
			if zero.notStyled {
				fmt.Println(str)
			} else {
				s := output.String(str).Foreground(output.Color("#ff0000")).Bold()
				fmt.Println(s)
			}
		}

		if zero.clearFunc == nil {
			Exit(1)
		} else {
			v := zero.clearFunc()
			path(v)
		}

	} else {
		if !zero.ifSucNotPrint {
			fmt.Println(R)
		}
	}

	path(zero.status)
}

func path(status programStatus) {
	switch status {
	case Brk:
		Exit(0)
	case Ftl:
		Exit(1)
	}
}
