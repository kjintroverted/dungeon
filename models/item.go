package models

type Item struct {
	Name     string  `json:"name"`
	Qty      int     `json:"qty"`
	Desc     string  `json:"description"`
	GoldCost float64 `json:"goldCost"`
}
