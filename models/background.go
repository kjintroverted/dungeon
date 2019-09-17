package models

import (
	"fmt"
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

func (r *Race) ParseTraits(rawTraits string) {
	fmt.Println(r.Name)
	traitStrings := strings.Split(rawTraits, "**_")
	for _, s := range traitStrings {
		vals := strings.Split(s, "._**")
		if len(vals) == 2 {
			r.Traits = append(r.Traits, TraitValues{Name: vals[0], Desc: strings.Trim(vals[1], " *\n")})
		} else if len(vals) == 1 {
			vals := strings.Split(s, "** ")
			if len(vals) == 2 {
				r.Traits = append(r.Traits, TraitValues{Name: strings.Trim(vals[0], " *\n"), Desc: strings.Trim(vals[1], " *\n")})
			}
		}
	}
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
