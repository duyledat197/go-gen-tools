package main

import (
	_ "net/http/pprof"

	"github.com/duyledat197/go-gen-tools/cmd"

	_ "github.com/jackc/pgx/v5"
)

func main() {
	cmd.Execute()
}
