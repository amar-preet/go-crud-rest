package main

import (
	"fmt"
	"go-crud-rest/db"
	"go-crud-rest/router"
	"log"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	r := router.Router(db)
	fmt.Println("Starting server on the port 8080...")
	r.Run("localhost:8080")
}
