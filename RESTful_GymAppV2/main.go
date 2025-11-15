package main

import (
	"log"
)

func main() {
	db := initDB()
	defer db.Close()

	// Ensure tables exist
	createUserInfoTable(db)
	createCoachInfoTable(db)

	// HTTP server
	router := setupRouter(db)
	log.Fatal(router.Run(":8080"))
}
