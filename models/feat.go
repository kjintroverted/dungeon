package models

// Feature models special traits and features for races and classes
type Feature struct {
	Index string   `json:"index,omitempty"`
	Name  string   `json:"name,omitempty"`
	Level int      `json:"level,omitempty"`
	URL   string   `json:"url,omitempty"`
	Desc  []string `json:"desc,omitempty"`
}
