package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/be9/protogobazel/lib/proto/coolpb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	// Protoc passes pluginpb.CodeGeneratorRequest in via stdin
	// marshalled with Protobuf
	input, _ := io.ReadAll(os.Stdin)
	var req pluginpb.CodeGeneratorRequest
	err := proto.Unmarshal(input, &req)
	if err != nil {
		panic(err)
	}

	// Initialise our plugin with default options
	opts := protogen.Options{}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	// Protoc passes a slice of File structs for us to process
	for _, file := range plugin.Files {
		// 1. Initialise a buffer to hold the generated code
		var buf bytes.Buffer
		empty := true

		// 2. Write the package name
		pkg := fmt.Sprintf("package %s", file.GoPackageName)
		buf.Write([]byte(pkg))

		// 3. For each message add our Foo() method
		for _, msg := range file.Proto.MessageType {
			ext := proto.GetExtension(msg.GetOptions(), coolpb.E_CoolMessage)
			if ext != nil {
				enabled := ext.(bool)
				if enabled {
					buf.Write([]byte(fmt.Sprintf(`
func (x %s) IsCool() bool {
   return true
}`, *msg.Name)))
					empty = false
				}
			}
		}

		// 4. Specify the output filename, in this case test.foo.go
		if !empty {
			filename := file.GeneratedFilenamePrefix + ".cool.go"
			f := plugin.NewGeneratedFile(filename, ".")

			// 5. Pass the data from our buffer to the plugin file struct
			_, err := f.Write(buf.Bytes())
			if err != nil {
				panic(err)
			}
		}
	}

	// Generate a response from our plugin and marshal as protobuf
	out, err := proto.Marshal(plugin.Response())
	if err != nil {
		panic(err)
	}

	// Write the response to stdout, to be picked up by protoc
	_, _ = fmt.Fprintf(os.Stdout, string(out))
}
