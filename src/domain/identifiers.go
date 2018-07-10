package domain

import (
	"encoding/json"
	"MA.Content.Services.OrgMapper/src/db"
)

type Identifiers struct {
	SourceInfo []string `json:"Source"`
	LabelInfo string `json:"Label"`
	ValueInfo string `json:"Value"`
}

func (i Identifiers) Get(key string) error {
	data, err := db.Get(key)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	return nil
}