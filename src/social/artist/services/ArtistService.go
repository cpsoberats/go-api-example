package services

import (
	"api/src/social/artist/entities"
	"api/src/social/artist/repositories"
)

type IArtistService interface {
	GetArtistList(from, size uint64) []entities.Artist
}

type ArtistService struct {
	ArtistRepository repositories.IArtistRepository
}

func (s ArtistService) GetArtistList(from, size uint64) []entities.Artist {
	artists := s.ArtistRepository.GetArtistList(from, size)
	return artists
}
