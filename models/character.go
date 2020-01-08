package models

import "github.com/kjintroverted/dungeon/util"

// Character models a playable character
type Character struct {
	ID              string    `json:"id" firestore:"-"`
	Owner           string    `json:"owner"`
	AuthorizedUsers []string  `json:"authUsers"`
	ReadUsers       []string  `json:"visibleTo"`
	Name            string    `json:"name"`
	Race            string    `json:"race"`
	Class           string    `json:"class"`
	Alignment       string    `json:"alignment"`
	XP              int       `json:"xp"`
	Level           int       `json:"level" firestore:"-"`
	ProBonus        int       `json:"proBonus" firestore:"-"`
	NextXP          int       `json:"next" firestore:"-"`
	HP              int       `json:"hp"`
	MaxHP           int       `json:"maxHP"`
	Armor           int       `json:"armor"`
	Speed           int       `json:"speed"`
	Str             int       `json:"str"`
	Dex             int       `json:"dex"`
	Con             int       `json:"con"`
	Intel           int       `json:"int"`
	Wis             int       `json:"wis"`
	Cha             int       `json:"cha"`
	Initiative      int       `json:"initiative,omitempty"`
	ProSkills       []string  `json:"proSkills"`
	ProTools        []string  `json:"proTools"`
	Languages       []string  `json:"languages"`
	Weapons         []Weapon  `json:"weapons"`
	Inventory       []Item    `json:"inventory"`
	GP              float64   `json:"gold"`
	Spells          []string  `json:"spells,omitempty"`
	Features        []string  `json:"feats,omitempty"`
	FeaturesNew     []Feature `json:"featsNew,omitempty"`
}

// PopulateLevelInfo will parse raw string data
// from the 5e SRD to fill in the character
func (c *Character) PopulateLevelInfo() {
	levelInfo := util.GetLevelInfo(c.XP)
	c.Level = levelInfo.Level
	c.NextXP = levelInfo.NextXP
	c.ProBonus = levelInfo.Bonus
}
