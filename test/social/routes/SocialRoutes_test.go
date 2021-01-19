package routes

import (
	socialRoutes "api/src/social/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	socialRoutes.Init(r.Group("/"))
	return r
}

func TestGetArtistStatus(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/social/artists/", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
