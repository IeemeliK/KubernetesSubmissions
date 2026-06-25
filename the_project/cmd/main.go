package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := "3000"
	if p := os.Getenv("PORT"); p != "" {
		PORT = p
	}

	addr := fmt.Sprintf(":%s", PORT)
	fmt.Printf("Server started in port %s\n", PORT)
	log.Fatal(http.ListenAndServe(addr, nil))
}
