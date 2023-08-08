package petstoreservice

import (
	"context"
	"fmt"
	"sync"
	"time"

	petv1 "buf.build/gen/go/acme/petapis/protocolbuffers/go/pet/v1"
	"connectrpc.com/connect"
	"github.com/oklog/ulid/v2"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/protobuf/types/known/durationpb"
)

type PetStoreService struct {
	sync.Mutex
	pets map[ulid.ULID]*petv1.Pet

	clock clock
}

func New() *PetStoreService {
	return &PetStoreService{
		pets:  map[ulid.ULID]*petv1.Pet{},
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
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}
	pet, ok := s.pets[petID]
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, nil)
	}
	return connect.NewResponse(&petv1.GetPetResponse{Pet: pet}), nil
}

func (s *PetStoreService) PutPet(
	ctx context.Context,
	req *connect.Request[petv1.PutPetRequest],
) (*connect.Response[petv1.PutPetResponse], error) {
	s.Lock()
	defer s.Unlock()
	createdAt := timeToDateTime(s.clock.Now())
	petID := ulid.Make()
	pet := &petv1.Pet{
		PetId:     petID.String(),
		PetType:   req.Msg.PetType,
		Name:      req.Msg.Name,
		CreatedAt: &createdAt,
	}
	s.pets[petID] = pet
	return connect.NewResponse(&petv1.PutPetResponse{Pet: pet}), nil
}

func (s *PetStoreService) DeletePet(
	ctx context.Context,
	req *connect.Request[petv1.DeletePetRequest],
) (*connect.Response[petv1.DeletePetResponse], error) {
	s.Lock()
	defer s.Unlock()
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

func (s *PetStoreService) PurchasePet(
	ctx context.Context,
	req *connect.Request[petv1.PurchasePetRequest],
) (*connect.Response[petv1.PurchasePetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("unimplemented"))
}

func timeToDateTime(t time.Time) datetime.DateTime {
	_, offset := t.Zone()
	return datetime.DateTime{
		Year:    int32(t.Year()),
		Month:   int32(t.Month()),
		Day:     int32(t.Day()),
		Hours:   int32(t.Hour()),
		Minutes: int32(t.Minute()),
		Seconds: int32(t.Second()),
		Nanos:   int32(t.Nanosecond()),
		TimeOffset: &datetime.DateTime_UtcOffset{
			UtcOffset: &durationpb.Duration{
				Seconds: int64(offset),
			},
		},
	}
}
