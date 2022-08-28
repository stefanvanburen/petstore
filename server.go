package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/bufbuild/connect-go"
	"github.com/oklog/ulid/v2"

	petv1 "github.com/stefanvanburen/petstore/gen/proto/go/pet/v1"
	"github.com/stefanvanburen/petstore/gen/proto/go/pet/v1/petv1connect"
)

type PetServer struct {
	sync.Mutex
	pets map[ulid.ULID]*petv1.Pet
}

func NewPetServer() *PetServer {
	return &PetServer{
		pets: map[ulid.ULID]*petv1.Pet{},
	}
}

func (s *PetServer) GetPet(
	ctx context.Context,
	req *connect.Request[petv1.GetPetRequest],
) (*connect.Response[petv1.GetPetResponse], error) {
	s.Lock()
	defer s.Unlock()
	petID, err := ulid.Parse(req.Msg.PetId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}
	if pet, ok := s.pets[petID]; ok {
		return connect.NewResponse(&petv1.GetPetResponse{Pet: pet}), nil
	}
	return nil, connect.NewError(connect.CodeNotFound, nil)
}

func (s *PetServer) PutPet(
	ctx context.Context,
	req *connect.Request[petv1.PutPetRequest],
) (*connect.Response[petv1.PutPetResponse], error) {
	s.Lock()
	defer s.Unlock()
	petID := ulid.Make()
	pet := &petv1.Pet{
		PetId:   petID.String(),
		PetType: req.Msg.PetType,
		Name:    req.Msg.Name,
		// TODO: CreatedAt
	}
	s.pets[petID] = pet
	return connect.NewResponse(&petv1.PutPetResponse{Pet: pet}), nil
}

func (s *PetServer) DeletePet(
	ctx context.Context,
	req *connect.Request[petv1.DeletePetRequest],
) (*connect.Response[petv1.DeletePetResponse], error) {
	petID, err := ulid.Parse(req.Msg.PetId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}
	if _, ok := s.pets[petID]; !ok {
		return nil, connect.NewError(connect.CodeNotFound, nil)
	}
	delete(s.pets, petID)
	return connect.NewResponse(&petv1.DeletePetResponse{}), nil
}

func (s *PetServer) PurchasePet(
	ctx context.Context,
	req *connect.Request[petv1.PurchasePetRequest],
) (*connect.Response[petv1.PurchasePetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("unimplemented"))
}

func main() {
	petServer := NewPetServer()
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(petServer)
	mux.Handle(path, handler)
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":8080", mux)
}
