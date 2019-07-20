package models

type Race struct {
	Name  string `json:"name"`
	Speed int    `json:"speed"`
}

type Class struct {
	Name string `json:"name"`
}
