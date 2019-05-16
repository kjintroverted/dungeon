package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kjintroverted/dungeon/api/srd"
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
	mux.HandleFunc("/api", srd.Root)
	mux.HandleFunc("/api/spells", srd.Spells)
	return
}

// ROOT GETTER
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Dungeon")
}
