package petstoreservice

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	petv1 "buf.build/gen/go/acme/petapis/protocolbuffers/go/pet/v1"
	"connectrpc.com/connect/v2"
	"go.vanburen.xyz/ok"
	"google.golang.org/genproto/googleapis/type/datetime"
)

func TestPetStoreService(t *testing.T) {
	synctest.Test(t, testPetStoreService)
}

// testPetStoreService runs in a synctest bubble, so time.Now is fixed until
// the test sleeps.
func testPetStoreService(t *testing.T) {
	petstoreservice := New()
	ctx := context.Background()
	putAt := time.Now()

	givenPet := &petv1.Pet{
		PetType: petv1.PetType_PET_TYPE_CAT,
		Name:    "Mobin",
	}

	putPetResponse, err := petstoreservice.PutPet(ctx, &petv1.PutPetRequest{
		PetType: givenPet.PetType,
		Name:    givenPet.Name,
	})
	ok.MustNoError(t, err)
	gotPutPet := putPetResponse.Pet
	ok.Equal(t, gotPutPet.Name, givenPet.Name)
	ok.Equal(t, gotPutPet.PetType, givenPet.PetType)
	ok.True(t, dateTimeToTime(gotPutPet.CreatedAt).Equal(putAt), ok.Sprintf("CreatedAt %v, want %v", gotPutPet.CreatedAt, putAt))

	petID := putPetResponse.Pet.PetId

	time.Sleep(time.Hour)

	getPetResponse, err := petstoreservice.GetPet(ctx, &petv1.GetPetRequest{
		PetId: petID,
	})
	ok.MustNoError(t, err)
	gotGetPet := getPetResponse.Pet
	ok.Equal(t, gotGetPet.Name, givenPet.Name)
	ok.Equal(t, gotGetPet.PetType, givenPet.PetType)
	ok.True(t, dateTimeToTime(gotGetPet.CreatedAt).Equal(putAt), ok.Sprintf("CreatedAt %v, want %v", gotGetPet.CreatedAt, putAt))

	_, err = petstoreservice.DeletePet(ctx, &petv1.DeletePetRequest{
		PetId: petID,
	})
	ok.MustNoError(t, err)

	_, err = petstoreservice.GetPet(ctx, &petv1.GetPetRequest{
		PetId: putPetResponse.Pet.PetId,
	})
	connectErr, isConnectErr := ok.ErrorAs[*connect.Error](t, err)
	if !isConnectErr {
		return
	}
	ok.Equal(t, connectErr.Code(), connect.CodeNotFound)
}

func dateTimeToTime(dt *datetime.DateTime) time.Time {
	zone := time.FixedZone("", int(dt.GetUtcOffset().GetSeconds()))
	return time.Date(int(dt.Year), time.Month(dt.Month), int(dt.Day), int(dt.Hours), int(dt.Minutes), int(dt.Seconds), int(dt.Nanos), zone)
}
