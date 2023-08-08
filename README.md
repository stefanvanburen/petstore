# petstore

This is a Go server that implements the [PetStoreService API](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService).
It's hosted on [fly.io](https://fly.io), at `petstore.fly.dev`.

It currently supports the following RPCs:

* [`PutPet`](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService.PutPet)
* [`GetPet`](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService.GetPet)
* [`DeletePet`](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService.DeletePet)

## Usage

You can interact with the API with plain HTTP requests (via the [Connect protocol](https://connect.build/docs/protocol/)) with any HTTP client, such as cURL, but
[`buf curl`](https://buf.build/docs/curl/usage/) makes it easy:

```console
$ buf curl --data '{"name": "Mobin", "petType": "PET_TYPE_CAT"}' --schema buf.build/acme/petapis https://petstore.fly.dev/pet.v1.PetStoreService/PutPet | jq .pet.petId # Create a pet
"01GT4XTKXEXY74QD8H575E8NWC"

$ buf curl --data '{"petId":"01GT4XTKXEXY74QD8H575E8NWC"}' --schema buf.build/acme/petapis https://petstore.fly.dev/pet.v1.PetStoreService/GetPet | jq .pet.name # Retrieve the pet
"Mobin"

$ buf curl --data '{"petId":"01GT4XTKXEXY74QD8H575E8NWC"}' --schema buf.build/acme/petapis https://petstore.fly.dev/pet.v1.PetStoreService/DeletePet # Delete a pet. :(
{}
```

You can also use [Buf Studio](https://studio.buf.build/acme/petapis/pet.v1.PetStoreService/PutPet?target=https%3A%2F%2Fpetstore.fly.dev) to interact with the API in a much more interactive way.

## Implementation details

The server uses the [connect-go library](https://github.com/connectrpc/connect-go) to implement the API.

The packages used for interacting with the API are [remotely generated](https://buf.build/docs/bsr/remote-packages/go/) - there's no code generation in this repository.

The "database" is completely in memory, so each deploy will wipe out any existing data.
