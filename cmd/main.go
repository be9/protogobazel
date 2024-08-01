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
		F: foopb.FooEnum_FOO_ENUM_KNOWN,
	}

	fmt.Println(prototext.Format(f))
	barpb.ExtraFn()
}
