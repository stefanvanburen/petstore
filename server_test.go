package main

import (
	"context"
	"errors"
	"testing"

	"github.com/akshayjshah/attest"
	"github.com/bufbuild/connect-go"
	petv1 "go.buf.build/bufbuild/connect-go/acme/petapis/pet/v1"
)

func TestServer(t *testing.T) {
	server := NewPetServer()
	ctx := context.Background()

	givenPet := &petv1.Pet{
		PetType: petv1.PetType_PET_TYPE_CAT,
		Name:    "Mobin",
	}

	putPetResponse, err := server.PutPet(ctx, connect.NewRequest(&petv1.PutPetRequest{
		PetType: givenPet.PetType,
		Name:    givenPet.Name,
	}))
	attest.Ok(t, err)
	attest.Equal(t, putPetResponse.Msg.Pet.Name, givenPet.Name)
	attest.Equal(t, putPetResponse.Msg.Pet.PetType, givenPet.PetType)

	petID := putPetResponse.Msg.Pet.PetId

	getPetResponse, err := server.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: petID,
	}))
	attest.Ok(t, err)
	attest.Equal(t, getPetResponse.Msg.Pet.Name, givenPet.Name)
	attest.Equal(t, getPetResponse.Msg.Pet.PetType, givenPet.PetType)

	_, err = server.DeletePet(ctx, connect.NewRequest(&petv1.DeletePetRequest{
		PetId: petID,
	}))
	attest.Ok(t, err)

	_, err = server.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: putPetResponse.Msg.Pet.PetId,
	}))
	var connectErr *connect.Error
	attest.True(t, errors.As(err, &connectErr))
	attest.Equal(t, connectErr.Code(), connect.CodeNotFound)
}
