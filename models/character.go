package models

import "github.com/kjintroverted/dungeon/util"

type Character struct {
	ID              string   `json:"id" firestore:"-"`
	Owner           string   `json:"owner"`
	AuthorizedUsers []string `json:"-"`
	Name            string   `json:"name"`
	Race            string   `json:"race"`
	Class           string   `json:"class"`
	XP              int      `json:"xp"`
	Level           int      `json:"level"`
	ProBonus        int      `json:"proBonus"`
	NextXP          int      `json:"next"`
	HP              int      `json:"hp"`
	Armor           int      `json:"armor"`
	Speed           int      `json:"speed"`
	Str             int      `json:"str"`
	Dex             int      `json:"dex"`
	Con             int      `json:"con"`
	Intel           int      `json:"intel"`
	Wis             int      `json:"wis"`
	Cha             int      `json:"cha"`
}

func (c *Character) PopulateLevelInfo() {
	levelInfo := util.GetLevelInfo(c.XP)
	c.Level = levelInfo.Level
	c.NextXP = levelInfo.NextXP
	c.ProBonus = levelInfo.Bonus
}
