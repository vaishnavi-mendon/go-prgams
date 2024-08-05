package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Define the route handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!") //Send "Hello WORLD!"as the response
	}
	//Create a new server mux
	mux := http.NewServeMux()
	//
	mux.HandleFunc("/", handler)
	//Crate a new HTTP server
	server := &http.Server{
		Addr:    ":8080", //Listen on port 8080
		Handler: mux,
	}
	//Start the server
	log.Println("Server listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}
