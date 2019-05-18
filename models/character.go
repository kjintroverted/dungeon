package models

type Character struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Race  string `json:"race"`
	Class string `json:"class"`
	Xp    int    `json:"xp"`
}
