package main

import (
	"context"
	"errors"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/matryer/is"
	petv1 "github.com/stefanvanburen/petstore/gen/proto/go/pet/v1"
)

func TestServer(t *testing.T) {
	is := is.New(t)
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
	is.NoErr(err)
	is.Equal(putPetResponse.Msg.Pet.Name, givenPet.Name)
	is.Equal(putPetResponse.Msg.Pet.PetType, givenPet.PetType)

	petID := putPetResponse.Msg.Pet.PetId

	getPetResponse, err := server.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: petID,
	}))
	is.NoErr(err)
	is.Equal(getPetResponse.Msg.Pet.Name, givenPet.Name)
	is.Equal(getPetResponse.Msg.Pet.PetType, givenPet.PetType)

	_, err = server.DeletePet(ctx, connect.NewRequest(&petv1.DeletePetRequest{
		PetId: petID,
	}))
	is.NoErr(err)

	_, err = server.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: putPetResponse.Msg.Pet.PetId,
	}))
	var connectErr *connect.Error
	is.True(errors.As(err, &connectErr))
	is.Equal(connectErr.Code(), connect.CodeNotFound)
}
