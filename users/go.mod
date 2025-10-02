module users

go 1.25.0

replace lib => ../lib

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.33.0-20240401165935-b983156c5e99.1
	github.com/bufbuild/protovalidate-go v0.6.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.2
	google.golang.org/genproto/googleapis/api v0.0.0-20250818200422-3122310a409c
	google.golang.org/grpc v1.75.1
	google.golang.org/protobuf v1.36.8
	lib v0.0.0-00010101000000-000000000000
)

require (
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/google/cel-go v0.20.1 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250826171959-ef028d996bc1 // indirect
)
