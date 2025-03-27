package main

import (
	"chatify-engine/router"
	"fmt"
)

func main() {
	r := router.Create()

	if err := r.Run(); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}
