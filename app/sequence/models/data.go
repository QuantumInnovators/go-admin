package models

type ClassList struct {
	Data []*ClassListKingdom `json:"data"`
}

type ClassListKingdom struct {
	Data   Kingdom            `json:"data"`
	Id     int                `json:"id"`
	Desc   string             `json:"label"`
	Phylum []*ClassListPhylum `json:"children"`
}

type ClassListPhylum struct {
	Data  Phylum            `json:"data"`
	Id    int               `json:"id"`
	Desc  string            `json:"label"`
	Class []*ClassListClass `json:"children"`
}

type ClassListClass struct {
	Data  Class             `json:"data"`
	Id    int               `json:"id"`
	Desc  string            `json:"label"`
	Order []*ClassListOrder `json:"children"`
}

type ClassListOrder struct {
	Data   Order              `json:"data"`
	Family []*ClassListFamily `json:"children"`
}

type ClassListFamily struct {
	Data  Family            `json:"data"`
	Genus []*ClassListGenus `json:"children"`
}

type ClassListGenus struct {
	Data    Genus               `json:"data"`
	Species []*ClassListSpecies `json:"children"`
}

type ClassListSpecies struct {
	Data Species `json:"data"`
}
