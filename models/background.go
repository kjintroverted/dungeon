package models

import (
	"strings"
)

type Race struct {
	Name  string `json:"name"`
	Speed int    `json:"speed"`
}

type Class struct {
	Name         string              `json:"name"`
	HitString    string              `json:"hit_dice"`
	SpellAbility string              `json:"spellcasting_ability,omitempty"`
	Information  map[string][]string `json:"info"`
	Table        string              `json:"table,omitempty" firestore:"-"`
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
