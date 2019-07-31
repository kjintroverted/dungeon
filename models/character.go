package models

import "github.com/kjintroverted/dungeon/util"

type Character struct {
	ID              string   `json:"id" firestore:"-"`
	Owner           string   `json:"owner"`
	AuthorizedUsers []string `json:"-"`
	Name            string   `json:"name"`
	Race            string   `json:"race"`
	Class           string   `json:"class"`
	HitDice         Dice     `json:"hitDice"`
	XP              int      `json:"xp"`
	Level           int      `json:"level" firestore:"-"`
	ProBonus        int      `json:"proBonus" firestore:"-"`
	NextXP          int      `json:"next" firestore:"-"`
	HP              int      `json:"hp"`
	Armor           int      `json:"armor"`
	Speed           int      `json:"speed"`
	Str             int      `json:"str"`
	Dex             int      `json:"dex"`
	Con             int      `json:"con"`
	Intel           int      `json:"intel"`
	Wis             int      `json:"wis"`
	Cha             int      `json:"cha"`
	Initiative      int      `json:"initiative,omitempty"`
	Spells          []string `json:"spells"`
}

func (c *Character) PopulateLevelInfo() {
	levelInfo := util.GetLevelInfo(c.XP)
	c.Level = levelInfo.Level
	c.NextXP = levelInfo.NextXP
	c.ProBonus = levelInfo.Bonus
}
