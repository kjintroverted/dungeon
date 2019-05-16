package util

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
