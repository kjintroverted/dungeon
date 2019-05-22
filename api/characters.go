package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/kjintroverted/dungeon/models"
	"google.golang.org/api/iterator"
)

func getAllCharacters(w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
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
		character.ID = doc.Ref.ID
		characters = append(characters, character)
	}

	res, _ := json.Marshal(characters)
	w.Write(res)
}

func updateCharacter(c models.Character, w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}

	result, _ := client.Collection("characters").Doc(c.ID).Set(ctx, c)

	bytes, _ := json.Marshal(result)
	w.Write(bytes)
}

func addCharacter(c models.Character, w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}

	ref, _, _ := client.Collection("characters").Add(ctx, c)

	fmt.Fprint(w, ref.ID)
}
