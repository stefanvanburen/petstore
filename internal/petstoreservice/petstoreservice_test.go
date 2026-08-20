package petstoreservice

import (
	"context"
	"testing"

	petv1 "buf.build/gen/go/acme/petapis/protocolbuffers/go/pet/v1"
	"connectrpc.com/connect"
	"go.vanburen.xyz/ok"
)

func TestPetStoreService(t *testing.T) {
	petstoreservice := New()
	ctx := context.Background()

	givenPet := &petv1.Pet{
		PetType: petv1.PetType_PET_TYPE_CAT,
		Name:    "Mobin",
	}

	putPetResponse, err := petstoreservice.PutPet(ctx, connect.NewRequest(&petv1.PutPetRequest{
		PetType: givenPet.PetType,
		Name:    givenPet.Name,
	}))
	ok.MustNoError(t, err)
	gotPutPet := putPetResponse.Msg.Pet
	ok.Equal(t, gotPutPet.Name, givenPet.Name)
	ok.Equal(t, gotPutPet.PetType, givenPet.PetType)

	petID := putPetResponse.Msg.Pet.PetId

	getPetResponse, err := petstoreservice.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: petID,
	}))
	ok.MustNoError(t, err)
	gotGetPet := getPetResponse.Msg.Pet
	ok.Equal(t, gotGetPet.Name, givenPet.Name)
	ok.Equal(t, gotGetPet.PetType, givenPet.PetType)

	_, err = petstoreservice.DeletePet(ctx, connect.NewRequest(&petv1.DeletePetRequest{
		PetId: petID,
	}))
	ok.MustNoError(t, err)

	_, err = petstoreservice.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: putPetResponse.Msg.Pet.PetId,
	}))
	connectErr, isConnectErr := ok.ErrorAs[*connect.Error](t, err)
	if !isConnectErr {
		return
	}
	ok.Equal(t, connectErr.Code(), connect.CodeNotFound)
}
