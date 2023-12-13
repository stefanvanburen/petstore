package petstoreservice

import (
	"context"
	"fmt"
	"sync"

	petv1 "buf.build/gen/go/acme/petapis/protocolbuffers/go/pet/v1"
	"connectrpc.com/connect"
	"github.com/oklog/ulid/v2"
)

type PetStoreService struct {
	sync.Mutex
	pets map[ulid.ULID]*pet

	clock clock
}

func New() *PetStoreService {
	return &PetStoreService{
		pets:  map[ulid.ULID]*pet{},
		clock: systemClock{},
	}
}

func (s *PetStoreService) GetPet(
	ctx context.Context,
	req *connect.Request[petv1.GetPetRequest],
) (*connect.Response[petv1.GetPetResponse], error) {
	s.Lock()
	defer s.Unlock()
	petID, err := ulid.Parse(req.Msg.PetId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("parsing pet id: %s", err))
	}
	pet, ok := s.pets[petID]
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("pet %q not found", petID))
	}
	return connect.NewResponse(&petv1.GetPetResponse{Pet: pet.ToProto()}), nil
}

func (s *PetStoreService) PutPet(
	ctx context.Context,
	req *connect.Request[petv1.PutPetRequest],
) (*connect.Response[petv1.PutPetResponse], error) {
	s.Lock()
	defer s.Unlock()
	pet := newPet(req.Msg.PetType, req.Msg.Name, s.clock.Now())
	s.pets[pet.id] = pet
	return connect.NewResponse(&petv1.PutPetResponse{Pet: pet.ToProto()}), nil
}

func (s *PetStoreService) DeletePet(
	ctx context.Context,
	req *connect.Request[petv1.DeletePetRequest],
) (*connect.Response[petv1.DeletePetResponse], error) {
	s.Lock()
	defer s.Unlock()
	petID, err := ulid.Parse(req.Msg.PetId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("parsing pet id: %s", err))
	}
	if _, ok := s.pets[petID]; !ok {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("pet %q not found", petID))
	}
	delete(s.pets, petID)
	return connect.NewResponse(&petv1.DeletePetResponse{}), nil
}

func (s *PetStoreService) PurchasePet(
	ctx context.Context,
	req *connect.Request[petv1.PurchasePetRequest],
) (*connect.Response[petv1.PurchasePetResponse], error) {
	s.Lock()
	defer s.Unlock()
	petID, err := ulid.Parse(req.Msg.PetId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("parsing pet id: %s", err))
	}
	if _, ok := s.pets[petID]; !ok {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("pet %q not found", petID))
	}
	delete(s.pets, petID)
	return connect.NewResponse(&petv1.PurchasePetResponse{}), nil
}
