package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kjintroverted/dungeon/api"
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
func createMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", root)

	r.HandleFunc("/api", srd.Root)
	r.HandleFunc("/api/spells", srd.Spells)

	r.HandleFunc("/api/characters", api.Characters)
	r.HandleFunc("/api/characters/{id}", api.Characters)
	return r
}

// ROOT GETTER
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Dungeon")
}
