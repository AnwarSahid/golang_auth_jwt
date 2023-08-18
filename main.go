package main

import (
	"golang_jwt/database"
	"golang_jwt/router"
)

func main() {
	database.StartDB()
	router.StartServer().Run(":8000")
}
