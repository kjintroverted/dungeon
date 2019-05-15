package srd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://api-beta.open5e.com/"

func Root(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(url)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(bytes))
}
