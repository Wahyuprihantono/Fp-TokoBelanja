package main

import (
	"Fp-TokoBelanja/database"
	"Fp-TokoBelanja/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	port := "8080"
	r.Run(":" + port)
}
