package main

import (
	"fmt"

	"advanced.microservices/pkg/store/postgres"
)

func main() {
	db, err := postgres.OpenDB("localhost", "", "jobbe", "password", "jobbe")
	if err != nil {
		fmt.Printf("Error connecting to database")
	}
	defer db.Close()
}
