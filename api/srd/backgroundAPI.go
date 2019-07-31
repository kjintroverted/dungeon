package srd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kjintroverted/dungeon/models"

	"github.com/kjintroverted/dungeon/util"
)

func GetRaces(w http.ResponseWriter, r *http.Request) {
	// REQUEST
	raw, err := util.Get(openURL + "/races")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// PULL RESULTS OFF RESPONSE
	var response interface{}
	json.Unmarshal(raw, &response)
	racesInterface := response.(map[string]interface{})["results"].([]interface{})

	// MAP RESULTS TO STRUCT
	var races []models.Race
	for _, i := range racesInterface {
		raceMap := i.(map[string]interface{})
		speed := raceMap["speed"].(map[string]interface{})
		races = append(races, models.Race{raceMap["name"].(string), int(speed["walk"].(float64))})
	}

	bytes, _ := json.Marshal(races)
	w.Write(bytes)
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	// REQUEST
	raw, err := util.Get(openURL + "/classes")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// PULL RESULTS OFF RESPONSE
	var response interface{}
	json.Unmarshal(raw, &response)
	classesInterface := response.(map[string]interface{})["results"].([]interface{})

	// MAP RESULTS TO STRUCT
	var classes []models.Class
	util.MapDecoder(&classes).Decode(classesInterface)

	bytes, _ := json.Marshal(classes)
	w.Write(bytes)
}
