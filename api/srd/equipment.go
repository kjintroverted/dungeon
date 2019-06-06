package srd

import (
	"encoding/json"
	"fmt"
	"net/http"

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
