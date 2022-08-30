# petstore

This is a Go server that implements the [PetStoreService API](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService).
It's hosted on <https://fly.io>, at `petstore.fly.dev`.

It currently supports the following RPCs:

* [`PutPet`](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService.PutPet)
* [`GetPet`](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService.GetPet)
* [`DeletePet`](https://buf.build/acme/petapis/docs/main:pet.v1#pet.v1.PetStoreService.DeletePet)

## Usage

```console
$ curl --header "Content-Type: application/json" --data '{"name": "Mobin", "petType": "PET_TYPE_CAT"}' https://petstore.fly.dev/pet.v1.PetStoreService/PutPet # Create a pet
{"pet":{"petType":"PET_TYPE_CAT", "petId":"01GBR5QYN85PVN8M8N0XFEJF15", "name":"Mobin"}}

$ curl --header "Content-Type: application/json" --data '{"petId": "01GBR5QYN85PVN8M8N0XFEJF15"}' https://petstore.fly.dev/pet.v1.PetStoreService/GetPet # Get a pet
{"pet":{"petType":"PET_TYPE_CAT", "petId":"01GBR5QYN85PVN8M8N0XFEJF15", "name":"Mobin"}}

$ curl --header "Content-Type: application/json" --data '{"petId": "01GBR5QYN85PVN8M8N0XFEJF15"}' https://petstore.fly.dev/pet.v1.PetStoreService/DeletePet # Delete a pet. :(
{}
```

## Implementation details

The server uses the [connect-go library](https://github.com/bufbuild/connect-go) to implement the API.

The "database" is completely in memory, so each deploy will wipe out any existing data.
