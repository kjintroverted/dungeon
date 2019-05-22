package util

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kjintroverted/dungeon/models"
	"github.com/mitchellh/mapstructure"
)

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

var advancement []models.Level = []models.Level{
	models.Level{0, 1, 2},
	models.Level{300, 2, 2},
	models.Level{900, 3, 2},
	models.Level{2700, 4, 2},
	models.Level{6500, 5, 3},
	models.Level{14000, 6, 3},
	models.Level{23000, 7, 3},
	models.Level{34000, 8, 3},
	models.Level{48000, 9, 4},
	models.Level{64000, 10, 4},
	models.Level{85000, 11, 4},
	models.Level{100000, 12, 4},
	models.Level{120000, 13, 5},
	models.Level{140000, 14, 5},
	models.Level{165000, 15, 5},
	models.Level{195000, 16, 5},
	models.Level{225000, 17, 6},
	models.Level{265000, 18, 6},
	models.Level{305000, 19, 6},
	models.Level{355000, 20, 6},
}
