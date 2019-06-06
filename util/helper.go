package util

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

var CoinConversion = map[string]float32{
	"cp": .01,
	"sp": .1,
	"ep": .5,
	"gp": 1,
	"pp": 10,
}

func MapDecoder(ref interface{}) *mapstructure.Decoder {
	config := mapstructure.DecoderConfig{TagName: "json", Result: ref}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return decoder
}

func Get(url string) ([]byte, error) {
	fmt.Println("GET", url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return bytes, nil
}

type Level struct {
	MinXP  int `json:"-"`
	Level  int `json:"level"`
	Bonus  int `json:"proBonus"`
	NextXP int `json:"next,omitempty"`
}

var Advancement = []Level{
	Level{MinXP: 0, Level: 1, Bonus: 2},
	Level{MinXP: 300, Level: 2, Bonus: 2},
	Level{MinXP: 900, Level: 3, Bonus: 2},
	Level{MinXP: 2700, Level: 4, Bonus: 2},
	Level{MinXP: 6500, Level: 5, Bonus: 3},
	Level{MinXP: 14000, Level: 6, Bonus: 3},
	Level{MinXP: 23000, Level: 7, Bonus: 3},
	Level{MinXP: 34000, Level: 8, Bonus: 3},
	Level{MinXP: 48000, Level: 9, Bonus: 4},
	Level{MinXP: 64000, Level: 10, Bonus: 4},
	Level{MinXP: 85000, Level: 11, Bonus: 4},
	Level{MinXP: 100000, Level: 12, Bonus: 4},
	Level{MinXP: 120000, Level: 13, Bonus: 5},
	Level{MinXP: 140000, Level: 14, Bonus: 5},
	Level{MinXP: 165000, Level: 15, Bonus: 5},
	Level{MinXP: 195000, Level: 16, Bonus: 5},
	Level{MinXP: 225000, Level: 17, Bonus: 6},
	Level{MinXP: 265000, Level: 18, Bonus: 6},
	Level{MinXP: 305000, Level: 19, Bonus: 6},
	Level{MinXP: 355000, Level: 20, Bonus: 6},
}
