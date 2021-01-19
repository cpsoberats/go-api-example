package artist

import (
	"api/src/social/artist/services"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func GetArtistService() services.IArtistService {
	return services.ArtistService{
		ArtistRepository: ArtistRepositoryMock{},
	}
}

func TestServiceGetAllArtist(t *testing.T) {
	service := GetArtistService()

	artistsList := service.GetArtistList(0, 10)

	assert.IsEqual(len(artistsList), 10)
}

