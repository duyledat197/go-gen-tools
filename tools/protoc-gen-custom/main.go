package main

import (
	//nolint
	"github.com/duyledat197/go-gen-tools/cmd/protoc-gen-custom/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	opt := protogen.Options{}
	internal.Run(opt)
}
