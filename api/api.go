package api

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"github.com/kjintroverted/dungeon/models"
	"github.com/kjintroverted/dungeon/util"
)

var app *firebase.App
var ctx context.Context
var client *firestore.Client
var err error

func Characters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllCharacters(w, r)
		break
	case "POST":
		// READ BODY BYTES
		raw, _ := ioutil.ReadAll(r.Body)

		// CONVERT JSON BYTES TO MAP
		var body map[string]interface{}
		json.Unmarshal(raw, &body)

		// CONVERT MAP TO STRUCT
		var character models.Character
		util.MapDecoder(&character).Decode(body)

		if len(character.ID) <= 0 {
			addCharacter(character, w, r)
		} else {
			updateCharacter(character, w, r)
		}
		break
	case "DELETE":
		deleteCharacter(mux.Vars(r)["id"], w, r)
		break
	default:
		e := errors.New("Invalid operation " + r.Method + " on character collection.")
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(e)
		w.Write(bytes)
	}
}
