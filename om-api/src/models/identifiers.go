package models

type IdentifiersObj struct {
	Count int `json:"count"`
	Identifiers []Identifier `json:"identifiers"` 
}

type Identifier struct {
	Source []string `json:"source"`
	Value string `json:"value"`
	Label string `json:"label"`
}