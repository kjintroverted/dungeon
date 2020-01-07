package models

type Feature struct {
	Index string   `json:"index"`
	Name  string   `json:"name"`
	Level int      `json:"level,omitempty"`
	URL   string   `json:"url"`
	Desc  []string `json:"desc,omitempty"`
}
