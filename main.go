package main

import (
	"fmt"

	"github.com/be9/protogobazel/lib/proto/barpb"
	"github.com/be9/protogobazel/lib/proto/foopb"
	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	b := &barpb.Bar{
		A: "aaaa",
	}
	f := &foopb.Foo{
		B: b,
	}

	fmt.Println(prototext.Format(f))
}
