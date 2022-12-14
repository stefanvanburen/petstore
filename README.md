# petstore

This is a Go server that implements the [PetStoreService API](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService).
It's hosted on [fly.io](https://fly.io), at `petstore.fly.dev`.

It currently supports the following RPCs:

* [`PutPet`](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService.PutPet)
* [`GetPet`](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService.GetPet)
* [`DeletePet`](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService.DeletePet)

## Usage

You can interact with the API with plain HTTP requests (via the [Connect protocol](https://connect.build/docs/protocol/)) with any HTTP client, such as cURL:

```console
$ curl --header "Content-Type: application/json" --data '{"name": "Mobin", "petType": "PET_TYPE_CAT"}' https://petstore.fly.dev/pet.v1.PetStoreService/PutPet # Create a pet
{"pet":{"petType":"PET_TYPE_CAT", "petId":"01GBR5QYN85PVN8M8N0XFEJF15", "name":"Mobin"}}

$ curl --header "Content-Type: application/json" --data '{"petId": "01GBR5QYN85PVN8M8N0XFEJF15"}' https://petstore.fly.dev/pet.v1.PetStoreService/GetPet # Get a pet
{"pet":{"petType":"PET_TYPE_CAT", "petId":"01GBR5QYN85PVN8M8N0XFEJF15", "name":"Mobin"}}

$ curl --header "Content-Type: application/json" --data '{"petId": "01GBR5QYN85PVN8M8N0XFEJF15"}' https://petstore.fly.dev/pet.v1.PetStoreService/DeletePet # Delete a pet. :(
{}
```

You can also use [Buf Studio](https://studio.buf.build/acme/petapis/pet.v1.PetStoreService/PutPet?target=https%3A%2F%2Fpetstore.fly.dev) to interact with the API in a much more interactive way.

## Implementation details

The server uses the [connect-go library](https://github.com/bufbuild/connect-go) to implement the API.

The "database" is completely in memory, so each deploy will wipe out any existing data.
