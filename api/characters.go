package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gorilla/websocket"
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

	owner := r.URL.Query().Get("owner")

	if owner != "" {
		iter = client.Collection("characters").Where("Owner", "==", owner).Documents(ctx)
	} else {
		iter = client.Collection("characters").Documents(ctx)
	}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var character models.Character
		doc.DataTo(&character)
		character.ID = doc.Ref.ID
		character.PopulateLevelInfo()
		characters = append(characters, character)
	}

	res, _ := json.Marshal(characters)
	w.Write(res)
}

func watchCharacters(w http.ResponseWriter, r *http.Request) {
	// DB CONNECTION
	ctx = context.Background()
	if app, err = firebase.NewApp(ctx, nil); err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}
	if client, err = app.Firestore(ctx); err != nil {
		fmt.Println("DB ERROR:", err.Error())
	}
	defer client.Close()

	// IDS TO WATCH
	idList := strings.Split(r.URL.Query().Get("id"), ",")

	// GET REF FOR EACH ENTITY FROM DB
	var docList []*firestore.DocumentRef
	for _, id := range idList {
		docList = append(docList, client.Collection("characters").Doc(id))
	}

	// SET UP WEBSOCKET
	var upgrader = websocket.Upgrader{}
	// ALLOW ALL ORIGINS
	upgrader.CheckOrigin = func(*http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// START A LOOP TO WAIT FOR A CLOSE MESSAGE FROM CLIENT
	go readLoop(c)

	// CREATE A CHANNEL FOR INCOMING DATA FROM DB
	snapChannel := make(chan *firestore.DocumentSnapshot)

	// GIVE EACH DOC A LOOP TO WAIT FOR DATA
	docs, _ := client.GetAll(ctx, docList)
	for _, doc := range docs {
		iter := doc.Ref.Snapshots(ctx)
		go dataLoop(iter, snapChannel)
	}

	// SEND DATA WHEN UPDATES COME THRU CHANNEL
	var characterSet []models.Character
	for {
		fmt.Println("Waiting for data")
		docSnap := <-snapChannel
		doc, _ := docSnap.Ref.Get(ctx)
		fmt.Println("Received data for", doc.Ref.ID)
		var character models.Character
		doc.DataTo(&character)
		character.ID = doc.Ref.ID
		character.PopulateLevelInfo()
		characterSet = addToCharacterSet(characterSet, character)

		bytes, _ := json.Marshal(characterSet)
		c.SetWriteDeadline(time.Now().Add(time.Second * 10))
		err := c.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			log.Println("Writing unavailable. Leaving Listener.")
			c.Close()
			break
		}
	}
}

func addToCharacterSet(characterSet []models.Character, character models.Character) []models.Character {
	for i, c := range characterSet {
		if c.ID == character.ID {
			result := append(characterSet[0:i], character)
			result = append(result, characterSet[i+1:]...)
			return result
		}
	}
	return append(characterSet, character)
}

func readLoop(c *websocket.Conn) {
	for {
		if _, _, err := c.NextReader(); err != nil {
			log.Println("Closing socket")
			c.Close()
			break
		}
	}
}

func dataLoop(iter *firestore.DocumentSnapshotIterator, c chan *firestore.DocumentSnapshot) {
	for {
		docsnap, _ := iter.Next()
		fmt.Println("Got data")
		c <- docsnap
	}
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

	docRef := client.Collection("characters").Doc(id)

	// GET CURRENT CHARACTER STATS
	doc, _ := docRef.Get(ctx)
	var character models.Character
	doc.DataTo(&character)
	character.ID = doc.Ref.ID
	character.PopulateLevelInfo()
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

	user := r.URL.Query().Get("user")

	var bytes []byte
	if len(user) > 0 {
		fmt.Println("checking:", user)
		type simpleResponse struct {
			Authorized bool `json:"authorized"`
		}
		result := user == character.Owner
		if result {
			bytes, _ = json.Marshal(simpleResponse{result})
			goto Send
		}
		for _, authUser := range character.AuthorizedUsers {
			if authUser == user {
				result = true
				break
			}
		}
		bytes, _ = json.Marshal(simpleResponse{result})
	} else {
		fmt.Println("getting all users")
		bytes, _ = json.Marshal(character.AuthorizedUsers)
	}
Send:
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
