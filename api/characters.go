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

func getAllCharacters(w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}
	defer client.Close()

	var characters []models.Character
	var iter *firestore.DocumentIterator
	if owner := r.URL.Query().Get("owner"); owner == "" {
		iter = client.Collection("characters").Documents(ctx)
	} else {
		iter = client.Collection("characters").Where("Owner", "==", owner).Documents(ctx)
	}

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

func getCharacter(id string, w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}
	defer client.Close()

	doc, _ := client.Collection("characters").Doc(id).Get(ctx)

	var character models.Character
	doc.DataTo(&character)
	character.ID = doc.Ref.ID
	bytes, _ := json.Marshal(character)
	w.Write(bytes)
}

func updateCharacter(c models.Character, w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}
	defer client.Close()

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
	defer client.Close()

	ref, _, _ := client.Collection("characters").Add(ctx, c)

	fmt.Fprint(w, ref.ID)
}

func getAuthUsers(id string, w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}
	defer client.Close()

	doc, _ := client.Collection("characters").Doc(id).Get(ctx)

	var character models.Character
	doc.DataTo(&character)

	bytes, _ := json.Marshal(character.AuthorizedUsers)

	w.Write(bytes)
}

func addAuthUser(id string, user string, w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}
	defer client.Close()

	doc, _ := client.Collection("characters").Doc(id).Get(ctx)

	var character models.Character
	doc.DataTo(&character)

	for _, email := range character.AuthorizedUsers {
		if email == user {
			w.Write([]byte(user + " already authorized."))
			return
		}
	}

	character.AuthorizedUsers = append(character.AuthorizedUsers, user)

	client.Collection("characters").Doc(doc.Ref.ID).Set(ctx, character)
	fmt.Println(character.AuthorizedUsers)

	w.Write([]byte("Added " + user + " as authorized user"))
}

func removeAuthUser(id string, user string, w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}
	defer client.Close()

	doc, _ := client.Collection("characters").Doc(id).Get(ctx)

	var character models.Character
	doc.DataTo(&character)

	arr := character.AuthorizedUsers
	for i, email := range arr {
		if email == user {
			character.AuthorizedUsers = append(arr[:i], arr[i+1:]...)
			break
		}
	}

	client.Collection("characters").Doc(doc.Ref.ID).Set(ctx, character)

	w.Write([]byte(user + " removed as authorized user"))
}

func deleteCharacter(id string, w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}
	defer client.Close()

	result, _ := client.Collection("characters").Doc(id).Delete(ctx)

	bytes, _ := json.Marshal(result)
	w.Write(bytes)
}
