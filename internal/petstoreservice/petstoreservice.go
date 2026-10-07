package petstoreservice

import (
	"context"
	"sync"
	"time"
	"uuid"

	petv1 "buf.build/gen/go/acme/petapis/protocolbuffers/go/pet/v1"
	"connectrpc.com/connect/v2"
)

type PetStoreService struct {
	sync.Mutex
	pets map[uuid.UUID]*pet
}

func New() *PetStoreService {
	return &PetStoreService{
		pets: map[uuid.UUID]*pet{},
	}
}

func (s *PetStoreService) GetPet(
	ctx context.Context,
	req *petv1.GetPetRequest,
) (*petv1.GetPetResponse, error) {
	s.Lock()
	defer s.Unlock()
	petID, err := uuid.Parse(req.PetId)
	if err != nil {
		return nil, connect.Errorf(connect.CodeInvalidArgument, "parsing pet id: %s", err)
	}
	pet, ok := s.pets[petID]
	if !ok {
		return nil, connect.Errorf(connect.CodeNotFound, "pet %q not found", petID)
	}
	return &petv1.GetPetResponse{Pet: pet.toProto()}, nil
}

func (s *PetStoreService) PutPet(
	ctx context.Context,
	req *petv1.PutPetRequest,
) (*petv1.PutPetResponse, error) {
	s.Lock()
	defer s.Unlock()
	pet := newPet(req.PetType, req.Name, time.Now())
	s.pets[pet.id] = pet
	return &petv1.PutPetResponse{Pet: pet.toProto()}, nil
}

func (s *PetStoreService) DeletePet(
	ctx context.Context,
	req *petv1.DeletePetRequest,
) (*petv1.DeletePetResponse, error) {
	s.Lock()
	defer s.Unlock()
	petID, err := uuid.Parse(req.PetId)
	if err != nil {
		return nil, connect.Errorf(connect.CodeInvalidArgument, "parsing pet id: %s", err)
	}
	if _, ok := s.pets[petID]; !ok {
		return nil, connect.Errorf(connect.CodeNotFound, "pet %q not found", petID)
	}
	delete(s.pets, petID)
	return &petv1.DeletePetResponse{}, nil
}

func (s *PetStoreService) PurchasePet(
	ctx context.Context,
	req *petv1.PurchasePetRequest,
) (*petv1.PurchasePetResponse, error) {
	s.Lock()
	defer s.Unlock()
	petID, err := uuid.Parse(req.PetId)
	if err != nil {
		return nil, connect.Errorf(connect.CodeInvalidArgument, "parsing pet id: %s", err)
	}
	if _, ok := s.pets[petID]; !ok {
		return nil, connect.Errorf(connect.CodeNotFound, "pet %q not found", petID)
	}
	delete(s.pets, petID)
	return &petv1.PurchasePetResponse{}, nil
}
