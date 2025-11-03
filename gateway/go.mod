module gateway

go 1.25.0

replace (
	auth => ../auth
	chat => ../chat
	lib => ../lib
	social => ../social
	users => ../users
)

require (
	auth v0.0.0-00010101000000-000000000000
	chat v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.2
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250826171959-ef028d996bc1
	google.golang.org/grpc v1.75.1
	social v0.0.0-00010101000000-000000000000
	users v0.0.0-00010101000000-000000000000
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.33.0-20240401165935-b983156c5e99.1 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250818200422-3122310a409c // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)
