package models

import (
	"errors"
	"strings"
)

type Race struct {
	Name   string        `json:"name"`
	Speed  int           `json:"speed"`
	Traits []TraitValues `json:"traits"`
}

type TraitValues struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Class struct {
	Name         string              `json:"name"`
	HitString    string              `json:"hit_dice"`
	SpellAbility string              `json:"spellcasting_ability,omitempty"`
	ProSaves     string              `json:"prof_saving_throws"`
	ProWeapons   string              `json:"prof_weapons"`
	Information  map[string][]string `json:"info"`
	Table        string              `json:"table,omitempty" firestore:"-"`
}

func (r *Race) ParseTraits(rawTraits string, visionString string) {
	// PARSE VISION STRINGS
	if trait, err := parseTraitString(visionString); err == nil {
		r.Traits = append(r.Traits, trait)
	}
	// PARSE TRAIT STRINGS
	traitStrings := strings.Split(rawTraits, "**_")
	for _, s := range traitStrings {
		if trait, err := parseTraitString(s); err == nil {
			r.Traits = append(r.Traits, trait)
		}
	}
}

func parseTraitString(s string) (TraitValues, error) {
	trimCharSet := " _*\n"
	if vals := strings.Split(s, "._**"); len(vals) == 2 {
		return TraitValues{Name: strings.Trim(vals[0], trimCharSet), Desc: strings.Trim(vals[1], trimCharSet)}, nil
	}
	if vals := strings.Split(s, "** "); len(vals) == 2 {
		return TraitValues{Name: strings.Trim(vals[0], trimCharSet), Desc: strings.Trim(vals[1], trimCharSet)}, nil
	}
	return TraitValues{}, errors.New("could not parse string")
}

func (c *Class) ParseTable() {
	rows := strings.Split(c.Table, "\n")
	var cols []string
	infoMap := make(map[string][]string)

	for _, data := range strings.Split(rows[0], "|") {
		value := strings.Trim(data, " -")
		infoMap[value] = []string{}
		cols = append(cols, value)
	}

	for _, r := range rows[2:] {
		cells := strings.Split(r, "|")
		for colNum, data := range cells {
			value := strings.Trim(data, " -")
			infoMap[cols[colNum]] = append(infoMap[cols[colNum]], value)
		}
	}

	delete(infoMap, "")
	c.Table = ""
	c.Information = infoMap
}
