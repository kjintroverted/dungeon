package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

// ALL CHARACTER HANDLERS
func Characters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if id := mux.Vars(r)["id"]; len(id) > 0 {
			getCharacter(id, w, r)
		} else {
			getAllCharacters(w, r)
		}
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

		fmt.Println("POST character", character.Name+"("+character.ID+")")

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

func AuthUsers(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	switch r.Method {
	case "GET":
		getAuthUsers(id, w, r)
		break
	case "POST":
		addAuthUser(id, r.URL.Query().Get("user"), w, r)
		break
	case "DELETE":
		removeAuthUser(id, r.URL.Query().Get("user"), w, r)
		break
	default:
		e := errors.New("Invalid operation " + r.Method + " on Auth Users collection.")
		w.WriteHeader(http.StatusBadRequest)
		bytes, _ := json.Marshal(e)
		w.Write(bytes)
	}
}

func LevelInfo(w http.ResponseWriter, r *http.Request) {
	xp, _ := strconv.Atoi(r.URL.Query().Get("xp"))

	var levelInfo util.Level
	for i, level := range util.Advancement {

		defer func() {
			if message := recover(); message != nil {
				bytes, _ := json.Marshal(level)
				w.Write(bytes)
			}
		}()

		nextLevel := util.Advancement[i+1]
		if nextLevel.MinXP > xp {
			levelInfo = level
			levelInfo.NextXP = nextLevel.MinXP
			break
		}
	}

	bytes, _ := json.Marshal(levelInfo)
	w.Write(bytes)
}
