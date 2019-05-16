package models

import (
	"strconv"
	"strings"
)

type Spell struct {
	Slug          string   `json:"slug"`
	Name          string   `json:"name"`
	Desc          string   `json:"desc"`
	Boosted       string   `json:"higher_level"`
	Range         string   `json:"range"`
	Components    string   `json:"components"`
	ComponentReq  []string `json:"componentReq"`
	Material      string   `json:"material"`
	Ritual        string   `json:"ritual"`
	Duration      string   `json:"duration"`
	Concentration string   `json:"concentration"`
	CastingTime   string   `json:"casting_time"`
	Level         string   `json:"level"`
	LevelNum      int      `json:"levelNum"`
	School        string   `json:"school"`
	Class         string   `json:"dnd_class"`
	Classes       []string `json:"classes"`
}

func (s *Spell) ConvertFields() {
	// SETS LEVEL NUM
	ordinal := strings.Split(s.Level, "-")[0]
	s.LevelNum, _ = strconv.Atoi(ordinal[:len(ordinal)-2])

	// SETS COMPONENT AND CLASS ARRAYS
	s.ComponentReq = strings.Split(s.Components, ", ")
	s.Classes = strings.Split(s.Class, ", ")
}
