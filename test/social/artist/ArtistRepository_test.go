package artist

import (
	"api/src/social/artist/repositories"
	"testing"
)

func GetArtistRepository() repositories.IArtistRepository {
	return ArtistRepositoryMock{}

}

func TestGetArtistsList(t *testing.T) {

	repository := GetArtistRepository()

	data := repository.GetArtistList(0, 10)

	if len(data) <= 0 {
		t.Fatalf("error")
	}
}

