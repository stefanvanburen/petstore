module github.com/stefanvanburen/petstore

go 1.22

toolchain go1.22.1

require (
	buf.build/gen/go/acme/petapis/connectrpc/go v1.16.2-20220907172654-7abdb7802c8f.1
	buf.build/gen/go/acme/petapis/protocolbuffers/go v1.34.2-20220907172654-7abdb7802c8f.2
	connectrpc.com/connect v1.16.2
	connectrpc.com/cors v0.1.0
	connectrpc.com/grpcreflect v1.2.0
	github.com/jba/templatecheck v0.7.0
	github.com/jub0bs/cors v0.2.0
	github.com/oklog/ulid/v2 v2.1.0
	go.akshayshah.org/attest v1.1.0
	golang.org/x/net v0.27.0
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6
	google.golang.org/protobuf v1.34.2
	rsc.io/markdown v0.0.0-20231030184305-7ce63bd70e80
)

require (
	buf.build/gen/go/acme/paymentapis/protocolbuffers/go v1.34.2-20220907172603-9a877cf260e1.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/safehtml v0.0.2 // indirect
	golang.org/x/text v0.16.0 // indirect
)
