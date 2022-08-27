package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"

	petv1 "github.com/stefanvanburen/petstore/gen/proto/go/pet/v1"
	"github.com/stefanvanburen/petstore/gen/proto/go/pet/v1/petv1connect"
)

type PetServer struct{}

func (s *PetServer) GetPet(
	ctx context.Context,
	req *connect.Request[petv1.GetPetRequest],
) (*connect.Response[petv1.GetPetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("unimplemented"))
}

func (s *PetServer) PutPet(
	ctx context.Context,
	req *connect.Request[petv1.PutPetRequest],
) (*connect.Response[petv1.PutPetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("unimplemented"))
}

func (s *PetServer) DeletePet(
	ctx context.Context,
	req *connect.Request[petv1.DeletePetRequest],
) (*connect.Response[petv1.DeletePetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("unimplemented"))
}

func (s *PetServer) PurchasePet(
	ctx context.Context,
	req *connect.Request[petv1.PurchasePetRequest],
) (*connect.Response[petv1.PurchasePetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("unimplemented"))
}

func main() {
	petServer := &PetServer{}
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(petServer)
	mux.Handle(path, handler)
	http.ListenAndServe(":8080", mux)
}
