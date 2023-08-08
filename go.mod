module github.com/stefanvanburen/petstore

go 1.19

require (
	buf.build/gen/go/acme/petapis/bufbuild/connect-go v1.10.0-20220907172654-7abdb7802c8f.1
	buf.build/gen/go/acme/petapis/protocolbuffers/go v1.31.0-20220907172654-7abdb7802c8f.1
	github.com/bufbuild/connect-go v1.10.0
	github.com/gomarkdown/markdown v0.0.0-20230322041520-c84983bdbf2a
	github.com/jub0bs/fcors v0.6.0
	github.com/oklog/ulid/v2 v2.1.0
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6
	google.golang.org/protobuf v1.31.0
)

require (
	buf.build/gen/go/acme/paymentapis/protocolbuffers/go v1.31.0-20220907172603-9a877cf260e1.1 // indirect
	golang.org/x/exp v0.0.0-20230801115018-d63ba01acd4b // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/text v0.11.0 // indirect
)
