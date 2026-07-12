package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	PORT := "3456"
	if p := os.Getenv("PORT"); p != "" {
		PORT = p
	}

	http.HandleFunc("/", homeHandler())

	addr := fmt.Sprintf(":%s", PORT)
	fmt.Printf("Server started in port %s\n", PORT)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func homeHandler() http.HandlerFunc {
	random := uuid.New()
	return func(w http.ResponseWriter, _ *http.Request) {
		_, err := fmt.Fprintf(w, "%s: %s", time.Now().Format("2006-01-02T15:04:05.000Z"), random)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
