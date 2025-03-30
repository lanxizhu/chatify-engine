package main

import (
	"chatify-engine/pkg/database"
	"chatify-engine/router"
	"fmt"
)

func main() {

	db, err := database.Create()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	r := router.Create(db)

	if err = r.Run(); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}
