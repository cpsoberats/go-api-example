package controllers

import (
	"api/src/social/artist/services"
)

type ArtistController struct {
	ArtistService services.IArtistService
}
