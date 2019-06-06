package srd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjintroverted/dungeon/models"
	"github.com/kjintroverted/dungeon/util"
)

const url string = "https://api-beta.open5e.com"

func GetSpells(w http.ResponseWriter, r *http.Request) {
	level := r.URL.Query().Get("level")
	var suffix string

	switch level {
	case "1":
		suffix = "st"
		break
	case "2":
		suffix = "nd"
		break
	case "3":
		suffix = "rd"
		break
	default:
		suffix = "th"
	}

	// REQUEST
	raw, err := util.Get(url + "/spells?level=" + level + suffix + "-level")
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

func GetSpell(w http.ResponseWriter, r *http.Request) {
	// REQUEST
	raw, err := util.Get(url + "/spells/" + mux.Vars(r)["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// PULL RESULTS OFF RESPONSE
	var response interface{}
	json.Unmarshal(raw, &response)
	spellInterface := response.(interface{})
	// MAP RESULTS TO STRUCT
	var spell models.Spell
	util.MapDecoder(&spell).Decode(spellInterface)
	spell.ConvertFields()

	// WRITE JSON
	b, _ := json.Marshal(spell)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}
	w.Write(b)
}
