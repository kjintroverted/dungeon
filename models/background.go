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
	for rowNum, r := range rows {
		cells := strings.Split(r, "|")
		colNum := -1
		for _, data := range cells {
			value := strings.Trim(data, " -")
			if value == "" {
				colNum++
				continue
			}
			if rowNum == 0 {
				infoMap[value] = []string{}
				cols = append(cols, value)
			} else {
				infoMap[cols[colNum]] = append(infoMap[cols[colNum]], value)
			}
			colNum++
		}
	}
	c.Table = ""
	c.Information = infoMap
}
