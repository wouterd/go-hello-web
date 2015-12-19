package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		port = "8080"
	}

	log.Printf("Starting a server on port %v", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
