package models

import (
	"strconv"
	"strings"
)

type Item struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url,omitempty"`
}

func (i *Item) ConvertFields() {
	params := strings.Split(i.URL, "/")
	i.Index, _ = strconv.Atoi(params[len(params)-1])
	i.URL = ""
}
