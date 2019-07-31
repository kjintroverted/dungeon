package models

type Weapon struct {
	Name       string   `json:"name"`
	Category   string   `json:"category"`
	DamageDice string   `json:"damage_dice"`
	DamageType string   `json:"damage_type"`
	Weight     string   `json:"weight"`
	Properties []string `json:"properties"`
}
