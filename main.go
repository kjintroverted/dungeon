package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	mux := createMux()
	fmt.Println("Server listening on", port)
	if error := http.ListenAndServe(":"+port, mux); error != nil {
		fmt.Println("ERROR", error)
	}
}

func createMux() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", root)
	return
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Dungeon")
}
