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
	raw, err := util.Get(openURL + "/spells?level=" + level + suffix + "-level")
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
