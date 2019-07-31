package srd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	if r.URL.Query().Get("name") != "" {
		GetClass(w, r)
		return
	}
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
	for i, c := range classes {
		c.ParseTable()
		classes[i] = c
	}

	bytes, _ := json.Marshal(classes)
	w.Write(bytes)
}

func GetClass(w http.ResponseWriter, r *http.Request) {
	// REQUEST
	raw, err := util.Get(openURL + "/classes/" + strings.ToLower(r.URL.Query().Get("name")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// PULL RESULTS OFF RESPONSE
	var response interface{}
	json.Unmarshal(raw, &response)

	// MAP RESULTS TO STRUCT
	var class models.Class
	util.MapDecoder(&class).Decode(response)
	class.ParseTable()
	bytes, _ := json.Marshal(class)
	w.Write(bytes)
}
