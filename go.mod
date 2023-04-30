module github.com/stefanvanburen/petstore

go 1.19

require (
	buf.build/gen/go/acme/petapis/bufbuild/connect-go v1.7.0-20220907172654-7abdb7802c8f.1
	buf.build/gen/go/acme/petapis/protocolbuffers/go v1.30.0-20220907172654-7abdb7802c8f.1
	github.com/bufbuild/connect-go v1.7.0
	github.com/jub0bs/fcors v0.2.0
	github.com/oklog/ulid/v2 v2.1.0
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6
	google.golang.org/protobuf v1.30.0
)

require (
	buf.build/gen/go/acme/paymentapis/protocolbuffers/go v1.30.0-20220907172603-9a877cf260e1.1 // indirect
	golang.org/x/exp v0.0.0-20230213192124-5e25df0256eb // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/text v0.7.0 // indirect
)
