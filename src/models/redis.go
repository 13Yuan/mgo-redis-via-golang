package models

type OrgRefrence struct {
	Mappings []MappingsInfo `json:"Mappings"`
}

type MappingsInfo struct {
	OrgRefrenceId string `json:"OrgRefrenceId"`
	InstRefrenceId string `json:"InstRefrenceId"`
	SaleRefrenceId string `json:"SaleRefrenceId"`
}

type OrgIdentifier struct {
	Cuisp []string `json:"Cusip"`
	Isin []string `json:"Isin"`
	Sedol []string `json:"Sedol"`
	Lei []string `json:"Lei"`
	Ticker []string `json:"Ticker"`
	Duns_Number []string `json:"Duns_Number"`
	BvD_ID []string `json:"BvD_ID"`
	ISIN_Nubmer []string `json:"ISIN_Nubmer"`
	Central_Index_Key []string `json:"Central_Index_Key"`
	Salesforce_Account_Id []string `json:"Salesforce_Account_Id"`
	SIC_Code []string `json:"SIC_Code"`
	Tokyo_Stock_Exchange_Ticker_Symbol []string `json:"Tokyo_Stock_Exchange_Ticker_Symbol"`
	Bank_Identifier_Code []string `json:"Bank_Identifier_Code"`
	Salesforce_Opportunity_Id []string `json:"Salesforce_Opportunity_Id"`
	Rssd_Id []string `json:"Rssd_Id"`
	Dtc_Sales_Agent_Part_Num []string `json:"Dtc_Sales_Agent_Part_Num"`
	Equity_Ticker []string `json:"Equity_Ticker"`
	Cmor_Company_Number []string `json:"Cmor_Company_Number"`
	Lloyds_Syndicate_Performance []string `json:"Lloyds_Syndicate_Performance"`
	Cu_Number []string `json:"Cu_Number"`
	Equity_SEDOL []string `json:"Equity_SEDOL"`
	Ibm_Number []string `json:"Ibm_Number"`
	FIGI []string `json:"FIGI"`
	FIGI_Previous []string `json:"FIGI_Previous"`
	FIGI_3 []string `json:"FIGI_3"`
	FIGI_4 []string `json:"FIGI_4"`
	FIGI_5 []string `json:"FIGI_5"`
	Moodys_Deal_Number []string `json:"Moodys_Deal_Number"`
	Legacy_Deal_Id []string `json:"Legacy_Deal_Id"`
}

type Identifier struct {
	OrgIdentifier
}