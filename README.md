# protogobazel

* `proto/src` contains 2 proto packages (`foo`, `bar`)
* `lib/proto` contains 2 Go packages with `.pb.go` generated from `proto/src`.
* `cmd/main.go` imports both packages and prints a debug Proto message.

## TODO

* regulate protoc versions
* generate `write_go_proto_sources` automatically
* add `compilers = [ ... ]` to `go_proto_library`
* figure out proto dep wtfs