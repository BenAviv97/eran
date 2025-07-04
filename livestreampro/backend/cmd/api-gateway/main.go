// /home/${USER}/livestreampro/backend/cmd/api-gateway/main.go
package main

import (
	"log"

	"livestreampro/internal/gateway"
)

func main() {
	e := gateway.NewRouter()

	log.Println("api-gateway started on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
