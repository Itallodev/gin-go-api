package main

import (
	"github.com/devitallo/gin-go-api/database"
	"github.com/devitallo/gin-go-api/routes"
)

func main() {
	database.ConectaDB()
	routes.HandleRequests()
}
