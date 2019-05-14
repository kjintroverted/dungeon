package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// SET PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// GET MULTIPLEX
	mux := createMux()

	// START SERVER
	fmt.Println("Server listening on", port)
	if error := http.ListenAndServe(":"+port, mux); error != nil {
		fmt.Println("ERROR", error)
	}
}

// CREATE MULTIPLEX PATH HANDLER
func createMux() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", root)
	return
}

// ROOT GETTER
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Dungeon")
}
