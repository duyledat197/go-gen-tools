package pathutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

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

func GetPkgDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}
