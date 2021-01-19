package controllers

import (
	"api/src/social/artist/entities"
)


type ArtistServiceMock struct {
}

func (s ArtistServiceMock) GetArtistList(from, size uint64) []entities.Artist {
	var artists []entities.Artist

	return artists
}
