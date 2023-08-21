package main

import (
	"fmt"
	"golang_jwt/database"
	"golang_jwt/router"
)

func main() {

	cek := "Sdasdsadsadas"

	cek2 := cek

	fmt.Println(cek2)

	database.StartDB()

	router.StartServer().Run(":8000")
}
