package models

type Item struct {
	Name     string `json:"name"`
	Desc     string `json:"description"`
	GoldCost float64    `json:"goldCost"`
}
