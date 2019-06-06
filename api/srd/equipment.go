package srd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjintroverted/dungeon/models"
	"github.com/kjintroverted/dungeon/util"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	raw, err := util.Get(dndURL + "/equipment")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// PULL RESULTS OFF RESPONSE
	var response interface{}
	json.Unmarshal(raw, &response)
	itemsInterface := response.(map[string]interface{})["results"].([]interface{})

	// MAP RESULTS TO STRUCT
	var items []models.Item
	util.MapDecoder(&items).Decode(itemsInterface)
	for i, item := range items {
		item.ConvertFields()
		items[i] = item
	}

	// WRITE JSON
	b, _ := json.Marshal(items)
	w.Write(b)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	// REQUEST
	raw, err := util.Get(dndURL + "/equipment/" + mux.Vars(r)["index"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// PULL RESULTS OFF RESPONSE
	var response interface{}
	json.Unmarshal(raw, &response)
	itemInterface := response.(interface{})

	// MAP RESULTS TO STRUCT
	var item models.Item
	util.MapDecoder(&item).Decode(itemInterface)
	item.ConvertFields()

	// WRITE JSON
	b, _ := json.Marshal(item)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}
	w.Write(b)
}
