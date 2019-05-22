package models

type Character struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Race  string `json:"race"`
	Class string `json:"class"`
	XP    int    `json:"xp"`
	Level int    `json:"level"`
	Str   int    `json:"str"`
	Dex   int    `json:"dex"`
	Con   int    `json:"con"`
	Intel int    `json:"intel"`
	Wis   int    `json:"wis"`
	Cha   int    `json:"cha"`
	HP    int    `json:"hp"`
	Armor int    `json:"armor"`
	Speed int    `json:"speed"`
}

type Level struct {
	XP               int
	Level            int
	ProficiencyBonus int
}
