package controllers

import (
	"api/src/core/apiconfig"
	"api/src/providers"
	"api/src/social/artist/http/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	providers.New(apiconfig.Service{})
}

func GetArtistController() controllers.ArtistController {
	return controllers.ArtistController{
		ArtistService: ArtistServiceMock{},
	}
}

func TestGetArtistList(t *testing.T) {

	ArtistController := controllers.ArtistController.GetAllHandler

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	req, _ := http.NewRequest("GET", "/api/v1/social/artists", nil)

	c.Request = req

	ArtistController(GetArtistController(), c)
}
