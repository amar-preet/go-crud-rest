package backend

import (
	"go-crud-rest/models"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=root dbname=User port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &gorm.DB{}, err
	}
	db.AutoMigrate(&models.Album{})
	return db, nil
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	var body []models.Album

	db, err := connectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	if result := db.Find(&body); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &body)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var body models.Album

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := connectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	if result := db.Create(&body); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.IndentedJSON(http.StatusCreated, &body)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var body models.Album
	db, err := connectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	if result := db.First(&body, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &body)
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var body models.Album
	db, err := connectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	if result := db.Delete(&body, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &body)
}

func UpdateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var body models.UpdateAlbumRequestBody

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db, err := connectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	var album models.Album

	if result := db.First(&album, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	album.Artist = body.Artist
	album.Title = body.Title
	album.Price = body.Price

	db.Save(&album)
	c.JSON(http.StatusOK, &album)
}
