package moduleutils

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/mod/modfile"
)

const (
	RED   = "\033[91m"
	RESET = "\033[0m"
)

func exitf(beforeExitFunc func(), code int, format string, args ...interface{}) {
	beforeExitFunc()
	fmt.Fprintf(os.Stderr, RED+format+RESET, args...)
	os.Exit(code)
}

func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		exitf(func() {}, 1, "%+v\n", err)
	}

	modName := modfile.ModulePath(goModBytes)

	return modName
}
