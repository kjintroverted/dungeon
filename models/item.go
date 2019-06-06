package models

import (
	"strconv"
	"strings"

	"github.com/kjintroverted/dungeon/util"
)

type Item struct {
	Index         int     `json:"index"`
	Name          string  `json:"name"`
	Category      string  `json:"equipment_category,omitempty"`
	WeaponRange   string  `json:"weapon_range,omitempty"`
	CategoryRange string  `json:"category_range,omitempty"`
	Damage        *Damage `json:"damage,omitempty"`
	Weight        int     `json:"weight,omitempty"`
	GoldCost      float32 `json:"goldCost,omitempty"`
	URL           string  `json:"url,omitempty"`
	Cost          *struct {
		Quantity int    `json:"quantity,omitempty"`
		Unit     string `json:"unit,omitempty"`
	} `json:"cost,omitempty"`
}

func (i *Item) ConvertFields() {
	params := strings.Split(i.URL, "/")
	i.Index, _ = strconv.Atoi(params[len(params)-1])
	i.URL = ""

	i.GoldCost = float32(i.Cost.Quantity) * util.CoinConversion[i.Cost.Unit]
	i.Cost = nil

	i.Damage.Type = i.Damage.DamageType.Name
	i.Damage.DamageType = nil
}
