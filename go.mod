module github.com/stefanvanburen/petstore

go 1.25.2

require (
	buf.build/gen/go/acme/petapis/connectrpc/go v1.19.1-20220907172654-7abdb7802c8f.2
	buf.build/gen/go/acme/petapis/protocolbuffers/go v1.36.11-20220907172654-7abdb7802c8f.1
	connectrpc.com/connect v1.19.1
	connectrpc.com/cors v0.1.0
	connectrpc.com/grpcreflect v1.3.0
	github.com/jba/templatecheck v0.7.1
	github.com/jub0bs/cors v0.9.2
	github.com/oklog/ulid/v2 v2.1.1
	go.akshayshah.org/attest v1.1.0
	google.golang.org/genproto v0.0.0-20241216192217-9240e9c98484
	google.golang.org/protobuf v1.36.11
	rsc.io/markdown v0.0.0-20241212154241-6bf72452917f
)

require (
	buf.build/gen/go/acme/paymentapis/protocolbuffers/go v1.36.11-20220907172603-9a877cf260e1.1 // indirect
	github.com/BurntSushi/toml v1.4.1-0.20240526193622-a339e1f7089c // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/safehtml v0.0.2 // indirect
	golang.org/x/exp/typeparams v0.0.0-20231108232855-2478ac86f678 // indirect
	golang.org/x/mod v0.28.0 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	golang.org/x/tools v0.37.0 // indirect
	golang.org/x/tools/go/expect v0.1.1-deprecated // indirect
	honnef.co/go/tools v0.6.1 // indirect
)

tool honnef.co/go/tools/cmd/staticcheck
