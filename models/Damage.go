package models

type Damage struct {
	DiceCount  int    `json:"dice_count"`
	DiceValue  int    `json:"dice_value"`
	Type       string `json:"type"`
	DamageType *struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"damage_type,omitempty"`
}
