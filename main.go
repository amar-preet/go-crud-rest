package main

import (
	"fmt"
	"go-crud-rest/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	r.Run("localhost:8080")
}
