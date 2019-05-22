package models

type Character struct {
	ID    string `json:"id" firestore:"-"`
	Name  string `json:"name"`
	Race  string `json:"race"`
	Class string `json:"class"`
	XP    int    `json:"xp"`
	HP    int    `json:"hp"`
	Armor int    `json:"armor"`
	Speed int    `json:"speed"`
	Str   int    `json:"str"`
	Dex   int    `json:"dex"`
	Con   int    `json:"con"`
	Intel int    `json:"intel"`
	Wis   int    `json:"wis"`
	Cha   int    `json:"cha"`
}

type Level struct {
	MinXP  int `json:"-"`
	Level  int `json:"level"`
	Bonus  int `json:"proBonus"`
	NextXP int `json:"next"`
}
