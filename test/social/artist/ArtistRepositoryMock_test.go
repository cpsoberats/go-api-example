package artist

import (
	"api/src/core/config/dataclient"
	"api/src/social/artist/entities"
	"api/test/core/config"
	"fmt"
)

type ArtistRepositoryMock struct {
}

func (r ArtistRepositoryMock) GetDataClient() dataclient.IDataClient {
	return config.DataClientMock{
		PostgresMockClient:        config.PostgresMockClient{},
	}
}

func (r ArtistRepositoryMock) GetArtistList(from, size uint64) []entities.Artist {
	var artists []entities.Artist

	for i:= 1; i < 20; i++{
		artist := entities.Artist{Id: int64(i), Name: fmt.Sprintf("Artist Test%d", i)}
		artists = append(artists, artist)
	}

	return artists
}

