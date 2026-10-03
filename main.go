package main

import (
	"demo-rest-api/db"
	"demo-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run("localhost:8080")

	if err != nil {
		return
	}
}
