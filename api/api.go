package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/kjintroverted/dungeon/models"
	"google.golang.org/api/iterator"
)

var app *firebase.App
var ctx context.Context
var client *firestore.Client
var err error

func Characters(w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	app, err = firebase.NewApp(ctx, nil)
	if err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}
	var characters []models.Character
	iter := client.Collection("characters").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var character models.Character
		doc.DataTo(&character)
		character.Id = doc.Ref.ID
		characters = append(characters, character)
	}

	res, _ := json.Marshal(characters)
	w.Write(res)
}
