package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (cnt ArtistController) GetAllHandler(c *gin.Context) {
	params := c.Request.URL.Query()
	size := 5
	from := 0

	if val, ok := params["size"]; ok {
		size, _ = strconv.Atoi(val[0])
	}
	if val, ok := params["from"]; ok {
		from, _ = strconv.Atoi(val[0])
	}

	artists := cnt.ArtistService.GetArtistList(uint64(from), uint64(size))

	c.JSON(http.StatusOK, artists)
}
