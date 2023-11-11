package petstoreservice

import (
	"context"
	"errors"
	"testing"

	petv1 "buf.build/gen/go/acme/petapis/protocolbuffers/go/pet/v1"
	"connectrpc.com/connect"
	"go.akshayshah.org/attest"
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
	attest.Ok(t, err)
	gotPutPet := putPetResponse.Msg.Pet
	attest.Equal(t, gotPutPet.Name, givenPet.Name)
	attest.Equal(t, gotPutPet.PetType, givenPet.PetType)

	petID := putPetResponse.Msg.Pet.PetId

	getPetResponse, err := petstoreservice.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: petID,
	}))
	attest.Ok(t, err)
	gotGetPet := getPetResponse.Msg.Pet
	attest.Equal(t, gotGetPet.Name, givenPet.Name)
	attest.Equal(t, gotGetPet.PetType, givenPet.PetType)

	_, err = petstoreservice.DeletePet(ctx, connect.NewRequest(&petv1.DeletePetRequest{
		PetId: petID,
	}))
	attest.Ok(t, err)

	_, err = petstoreservice.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: putPetResponse.Msg.Pet.PetId,
	}))
	var connectErr *connect.Error
	isConnectErr := errors.As(err, &connectErr)
	attest.True(t, isConnectErr)
	attest.Equal(t, connectErr.Code(), connect.CodeNotFound)
}
