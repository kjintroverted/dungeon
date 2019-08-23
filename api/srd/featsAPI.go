package srd

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kjintroverted/dungeon/models"
	"github.com/kjintroverted/dungeon/util"
)

func GetSpells(w http.ResponseWriter, r *http.Request) {
	var spells []models.Spell

	level := r.URL.Query().Get("level")
	if level != "" {
		spells, _ = getSpellsForLevel(level)
	}

	slugs := r.URL.Query().Get("name")
	if slugs != "" {
		spells, _ = getSpellsBySlug(strings.Split(slugs, ","))
	}

	// WRITE JSON
	b, _ := json.Marshal(spells)
	w.Write(b)
}

func getSpellsForLevel(level string) ([]models.Spell, error) {
	if level != "Cantrip" {
		level += "-level"
	}

	// REQUEST
	raw, err := util.Get(openURL + "/spells?level=" + level)
	if err != nil {
		return nil, err
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
	return spells, nil
}

func getSpellsBySlug(slugs []string) ([]models.Spell, error) {
	var spells []models.Spell
	for _, slug := range slugs {
		raw, err := util.Get(openURL + "/spells/" + slug)
		if err != nil {
			return spells, err
		}

		// PULL RESULTS OFF RESPONSE
		var response interface{}
		json.Unmarshal(raw, &response)
		// MAP RESULTS TO STRUCT
		var spell models.Spell
		util.MapDecoder(&spell).Decode(response.(interface{}))
		spell.ConvertFields()
		spells = append(spells, spell)
	}
	return spells, nil
}

func GetFeats(w http.ResponseWriter, r *http.Request) {
	var feats []models.Feature

	ids := r.URL.Query().Get("id")
	if ids != "" {
		feats = getFeatsByID(strings.Split(ids, ","))
	} else {
		feats = allFeats()
	}

	// WRITE JSON
	b, _ := json.Marshal(feats)
	w.Write(b)
}

func allFeats() []models.Feature {
	raw, _ := util.Get(dndURL + "/features")

	// PULL RESULTS OFF RESPONSE
	var response interface{}
	json.Unmarshal(raw, &response)

	featsInterface := response.(map[string]interface{})["results"].([]interface{})
	// MAP RESULTS TO STRUCT
	var feats []models.Feature
	util.MapDecoder(&feats).Decode(featsInterface)

	return uniqueFeats(feats)
}

func uniqueFeats(arr []models.Feature) (result []models.Feature) {
	for i, feat := range arr {
		if i%3 == 0 {
			result = append(result, feat)
		}
	}
	return
}

func getFeatsByID(ids []string) []models.Feature {
	var feats []models.Feature
	for _, id := range ids {
		raw, _ := util.Get(dndURL + "/features/" + id)

		// PULL RESULTS OFF RESPONSE
		var response interface{}
		json.Unmarshal(raw, &response)
		// MAP RESULTS TO STRUCT
		var feature models.Feature
		util.MapDecoder(&feature).Decode(response.(interface{}))
		feats = append(feats, feature)
	}
	return feats
}
