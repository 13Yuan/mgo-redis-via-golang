package models

type DocsInfo struct {
	Lei []string `json:"org_lei"`
	Cusip []string  `json:"org_Cusip"`
	Isin []string `json:"org_isin"`
	Sedol []string `json:"org_sedol"`
}

type DocListInfo struct {
	NumFound int `json:"numFound"`
	Start int `json:"start"`
	Docs []DocsInfo `json:"docs"`
}

type Group struct {
	GroupValue string `json:"groupValue"`
	Doclist DocListInfo `json:"doclist"`
}

type OdsRefIdGroupInfo struct {
	Matches int `json:"matches"`
	Groups []Group `json:"groups"`
}

type GroupedInfo struct {
	OdsRefIdGroup OdsRefIdGroupInfo `json:"ods_ref_id_group"`
}

type OrgMapping struct {
	Grouped GroupedInfo `json:"grouped"`
}