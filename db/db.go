package db

import (
	"go-crud-rest/models"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get("POSTGRES_URL").(string)
	if !ok {
		log.Fatalf("Invalid Database connection string")
	}

	db, err := gorm.Open(postgres.Open(value), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &gorm.DB{}, err
	}

	db.AutoMigrate(&models.Album{})
	return db, nil
}
