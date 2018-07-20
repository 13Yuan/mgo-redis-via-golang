package models

type (
	IDs struct {
		ID string `json:"_id" bson:"_id"`
		OrgID int `json:"org_id" bson:"org_id"`
		Identifiers []IdentifiersObj `json:"identifiers" bson:"identifiers"`
	}
	Sale struct {
		ID string `bson:"_id"`
		InstId int `bson:"inst_id"`
		OrgIds []int `bson:"org_id"`
		SaleOrgs []SaleOrg `bson:"sale_org"`
	}
	SaleOrg struct {
		OrgId int `json:"org_id" bson:"org_id"`
		OrganizationRole string `json:"organization_role" bson:"organization_role"`
	}
	IdentifiersObj struct {
		Source []string `json:"source" bson:"source"`
		Label string `json:"label" bson:"label"`
		Value string `json:"value" bson:"value"`	
	}
	KeyValue struct {
		Key string
		Value []byte
	}
	KeysObj struct {
		Keys []string
		Id IDs
	}
)