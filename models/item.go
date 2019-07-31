package models

type Item struct {
	Name     string `json:"name"`
	Desc     string `json:"description"`
	GoldCost int    `json:"goldCost"`
}
