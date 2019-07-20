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
	r.HandleFunc("/api", root)

	r.HandleFunc("/api/races", srd.GetRaces)
	r.HandleFunc("/api/classes", srd.GetClasses)
	r.HandleFunc("/api/spells", srd.GetSpells)
	r.HandleFunc("/api/spells/{name}", srd.GetSpell)
	r.HandleFunc("/api/items", srd.GetItems)
	r.HandleFunc("/api/items/{index}", srd.GetItem)

	r.HandleFunc("/api/characters", api.Characters)
	r.HandleFunc("/api/characters/{id}", api.Characters)
	r.HandleFunc("/api/characters/{id}/auth-users", api.AuthUsers)
	r.HandleFunc("/api/level", api.LevelInfo)

	return r
}

// ROOT GETTER
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Dungeon")
}
