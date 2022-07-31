package router

import (
	"go-crud-rest/backend"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	h := &backend.Handler{
		DB: db,
	}

	router := gin.Default()
	router.GET("/albums", h.GetAlbums)
	router.GET("/albums/:id", h.GetAlbumByID)
	router.POST("/albums", h.PostAlbums)
	router.DELETE("/albums/:id", h.DeleteAlbumByID)
	router.PUT("/albums/:id", h.UpdateAlbumByID)

	return router
}
