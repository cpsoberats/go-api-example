package artist

import (
	//coreConfig "api/src/core/config"
	"api/src/social/artist/config"
	"api/src/social/artist/http/controllers"
	"github.com/gin-gonic/gin"
)

func ArtistController() controllers.ArtistController {
	return config.
		Container().
		Get("artist_controller").(controllers.ArtistController)
}

func Routes(socialGroup *gin.RouterGroup) {

	// Artists list
	socialGroup.GET("/artists",
		ArtistController().GetAllHandler)

}