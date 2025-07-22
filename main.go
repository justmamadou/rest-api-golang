package main

import (
	"github.com/gin-gonic/gin"
	"github.com/justmamadou/rest-api-golang/db"
	"github.com/justmamadou/rest-api-golang/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
