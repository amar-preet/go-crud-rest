package backend

import (
	"go-crud-rest/models"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

// getAlbums
func (h Handler) GetAlbums(c *gin.Context) {
	var body []models.Album

	if result := h.DB.Find(&body); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &body)
}

// postAlbums
func (h Handler) PostAlbums(c *gin.Context) {
	var body models.Album
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := h.DB.Create(&body); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.IndentedJSON(http.StatusCreated, &body)
}

// get album by ID
func (h Handler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var body models.Album
	if result := h.DB.First(&body, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &body)
}

// delete album by ID
func (h Handler) DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var body models.Album
	if result := h.DB.Delete(&body, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &body)
}

// update album by ID
func (h Handler) UpdateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var body models.UpdateAlbumRequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var album models.Album
	if result := h.DB.First(&album, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	album.Artist = body.Artist
	album.Title = body.Title
	album.Price = body.Price

	h.DB.Save(&album)
	c.JSON(http.StatusOK, &album)
}
