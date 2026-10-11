package utils

import (
	interpol "github.com/imkira/go-interpol"
	"os"
	"strings"
)

func Exit(code int) {
	os.Exit(code)
}

func Args() []string {
	return os.Args
}

func StartWithError(o error, t string) bool {
	return strings.HasPrefix(o.Error(), t)
}

type ItpMap map[string]string

func (m ItpMap) BuildStrWithInterpol(templ string) (string, error) {
	str, err := interpol.WithMap(templ, m)
	if err != nil {
		return "", err
	}
	return str, nil
}
