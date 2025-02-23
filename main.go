package main

import (
	"log"

	"github.com/umbe77/my-ado-prs/database"
)

func main() {
	err := database.Migrate()
	if err != nil {
		log.Println("Error open db", err)
	}

	database.Close()
}
