package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := "3000"
	if p := os.Getenv("PORT"); p != "" {
		PORT = p
	}

	http.HandleFunc("/", homeHandler)

	addr := fmt.Sprintf(":%s", PORT)
	fmt.Printf("Server started in port %s\n", PORT)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h1>Hello World</h1>
		</body>
		</html>	
	`
	tmpl := template.New("index")
	tmpl.Parse(html)

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("error parsing html %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
