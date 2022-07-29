package router

import (
	"go-crud-rest/backend"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", backend.GetAlbums)
	router.GET("/albums/:id", backend.GetAlbumByID)
	router.POST("/albums", backend.PostAlbums)
	router.DELETE("/albums/:id", backend.DeleteAlbumByID)
	router.PUT("/albums/:id", backend.UpdateAlbumByID)

	return router
}
