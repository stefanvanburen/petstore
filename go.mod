module github.com/stefanvanburen/petstore

go 1.19

require (
	buf.build/gen/go/acme/petapis/connectrpc/go v1.11.1-20220907172654-7abdb7802c8f.1
	buf.build/gen/go/acme/petapis/protocolbuffers/go v1.31.0-20220907172654-7abdb7802c8f.1
	connectrpc.com/connect v1.11.1
	connectrpc.com/grpcreflect v1.2.0
	github.com/gomarkdown/markdown v0.0.0-20230922112808-5421fefb8386
	github.com/jub0bs/fcors v0.6.0
	github.com/oklog/ulid/v2 v2.1.0
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6
	google.golang.org/protobuf v1.31.0
)

require (
	buf.build/gen/go/acme/paymentapis/protocolbuffers/go v1.31.0-20220907172603-9a877cf260e1.1 // indirect
	golang.org/x/exp v0.0.0-20230801115018-d63ba01acd4b // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
