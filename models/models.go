package models

import "gorm.io/gorm"

// album represents data about a record album.
type Album struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Price  int    `json:"price" binding:"required"`
}

type UpdateAlbumRequestBody struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  int    `json:"price"`
}
