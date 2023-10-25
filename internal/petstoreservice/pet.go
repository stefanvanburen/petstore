package petstoreservice

import (
	"time"

	petv1 "buf.build/gen/go/acme/petapis/protocolbuffers/go/pet/v1"
	"github.com/oklog/ulid/v2"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/protobuf/types/known/durationpb"
)

type pet struct {
	id        ulid.ULID
	typ       petv1.PetType
	name      string
	createdAt time.Time
}

func newPet(petType petv1.PetType, name string, createdAt time.Time) *pet {
	petID := ulid.Make()
	return &pet{
		id:        petID,
		typ:       petType,
		name:      name,
		createdAt: createdAt,
	}
}

func (p *pet) ToProto() *petv1.Pet {
	return &petv1.Pet{
		PetId:     p.id.String(),
		PetType:   p.typ,
		Name:      p.name,
		CreatedAt: timeToDateTime(p.createdAt),
	}
}

func timeToDateTime(t time.Time) *datetime.DateTime {
	_, offset := t.Zone()
	return &datetime.DateTime{
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
