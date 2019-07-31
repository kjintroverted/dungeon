package srd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kjintroverted/dungeon/models"
	"github.com/kjintroverted/dungeon/util"
)

func GetWeapons(w http.ResponseWriter, r *http.Request) {
	// REQUEST
	raw, err := util.Get(openURL + "/weapons/")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// PULL RESULTS OFF RESPONSE
	var response interface{}
	json.Unmarshal(raw, &response)
	weaponInterface := response.(map[string]interface{})["results"].([]interface{})
	// MAP RESULTS TO STRUCT
	var weapons []models.Weapon
	util.MapDecoder(&weapons).Decode(weaponInterface)

	// WRITE JSON
	b, _ := json.Marshal(weapons)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}
	w.Write(b)
}
