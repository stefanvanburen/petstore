package server

import (
	"context"
	"errors"
	"testing"

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
	if err != nil {
		t.Errorf("PutPet: got %v, want err = nil", err)
	}
	gotPutPet := putPetResponse.Msg.Pet
	if gotPutPet.Name != givenPet.Name {
		t.Errorf("PutPet: got %s name, want %s name", gotPutPet.Name, givenPet.Name)
	}
	if gotPutPet.PetType != givenPet.PetType {
		t.Errorf("PutPet: got %v type, want %v type", gotPutPet.PetType, givenPet.PetType)
	}

	petID := putPetResponse.Msg.Pet.PetId

	getPetResponse, err := server.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: petID,
	}))
	if err != nil {
		t.Errorf("GetPet: got %v, want err = nil", err)
	}
	gotGetPet := getPetResponse.Msg.Pet
	if gotGetPet.Name != givenPet.Name {
		t.Errorf("GetPet: got %s name, want %s name", gotGetPet.Name, givenPet.Name)
	}
	if gotGetPet.PetType != givenPet.PetType {
		t.Errorf("GetPet: got %v type, want %v type", gotGetPet.PetType, givenPet.PetType)
	}

	_, err = server.DeletePet(ctx, connect.NewRequest(&petv1.DeletePetRequest{
		PetId: petID,
	}))
	if err != nil {
		t.Errorf("DeletePet: got %v, want err = nil", err)
	}

	_, err = server.GetPet(ctx, connect.NewRequest(&petv1.GetPetRequest{
		PetId: putPetResponse.Msg.Pet.PetId,
	}))
	var connectErr *connect.Error
	if isConnectErr := errors.As(err, &connectErr); !isConnectErr {
		t.Errorf("GetPet: got %v, want connectErr", err)
	}
	if connectErr.Code() != connect.CodeNotFound {
		t.Errorf("GetPet: got %v code, want connect.CodeNotFound", connectErr.Code())
	}
}
