package models

type OrgRefrence struct {
	Mappings []MappingsInfo `json:"Mappings"`
}

type MappingsInfo struct {
	OrgRefrenceId string `json:"OrgRefrenceId"`
}

type Identifier struct {
	Cuisp []string `json:"Cusip"`
	Isin []string `json:"Isin"`
	Sedol []string `json:"Sedol"`
	Lei []string `json:"Lei"`
}