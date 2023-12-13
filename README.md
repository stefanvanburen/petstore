# PetStore 🐶🐱🏪

This is a Go server that implements the [PetStoreService API](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService).
It's hosted on [fly.io](https://fly.io), at [petstore.fly.dev](https://petstore.fly.dev).
The source code is at [github.com/stefanvanburen/petstore](https://github.com/stefanvanburen/petstore).

## Usage

You can interact with the API with plain HTTP requests (via the [Connect protocol](https://connectrpc.com/docs/protocol/)) with any HTTP client, such as cURL, but
[`buf curl`](https://buf.build/docs/curl/usage/) makes it easy:

```console
$ # Create a pet
$ buf curl \
  --data '{"name": "Mobin", "petType": "PET_TYPE_CAT"}' \
  https://petstore.fly.dev/pet.v1.PetStoreService/PutPet | jq .pet.petId
"01GT4XTKXEXY74QD8H575E8NWC"

$ # Retrieve a pet
$ buf curl \
  --data '{"petId":"01GT4XTKXEXY74QD8H575E8NWC"}' \
  https://petstore.fly.dev/pet.v1.PetStoreService/GetPet | jq .pet.name
"Mobin"

$ # Delete a pet. :(
$ buf curl \
  --data '{"petId":"01GT4XTKXEXY74QD8H575E8NWC"}' \
  https://petstore.fly.dev/pet.v1.PetStoreService/DeletePet
{}
```

You can also use [Buf Studio](https://buf.build/studio/acme/petapis/pet.v1.PetStoreService/PutPet?target=https%3A%2F%2Fpetstore.fly.dev) to interact with the API in a much more interactive way.

## Implementation details

The server uses the [connect-go library](https://github.com/connectrpc/connect-go) to implement the API, with [connectrpc/grpcreflect-go](https://github.com/connectrpc/grpcreflect-go) adding support for the gRPC server reflection API.

The packages used for interacting with the API are [remotely generated](https://buf.build/docs/bsr/remote-packages/go/) - there's no code generation in this repository.

The "database" is completely in memory, so each deploy will wipe out any existing data.
