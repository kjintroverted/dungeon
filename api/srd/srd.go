package srd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kjintroverted/dungeon/models"
	"github.com/kjintroverted/dungeon/util"
)

const url string = "https://api-beta.open5e.com"

func Spells(w http.ResponseWriter, r *http.Request) {
	// REQUEST
	raw, err := util.Get(url + "/spells?" + r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// PULL RESULTS OFF RESPONSE
	var response interface{}
	json.Unmarshal(raw, &response)
	spellsInterface := response.(map[string]interface{})["results"].([]interface{})

	// MAP RESULTS TO STRUCT
	var spells []models.Spell
	util.MapDecoder(&spells).Decode(spellsInterface)
	for i, spell := range spells {
		spell.ConvertFields()
		spells[i] = spell
	}

	// WRITE JSON
	b, _ := json.Marshal(spells)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}
	w.Write(b)
}
