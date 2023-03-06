package main

import (
	"github.com/duyledat197/go-gen-tools/tools/protoc-gen-custom/internal"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	opt := protogen.Options{}
	internal.Run(opt)
}
