package routes

import (
	"api/src/social/artist"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.RouterGroup) *gin.RouterGroup {
	socialGroup := r.Group("/social")

	socialGroup.GET("/artists/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "active",
		})
	})

	artist.Routes(socialGroup)

	return r
}
