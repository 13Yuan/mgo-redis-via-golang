package models

func GetOrgReferencePrefix() []string {
	prefix := []string {
		"cusip",
		"lei",
		"duns_number",
		"bvd_id",
		"isin_number",
		"central_index_key",
		"isin",
		"salesforce_account_id",
		"sedol",
		"ticker",
		"bloomberg_id",
		"sic_code",
		"tokyo_stock_exchange_ticker_symbol",
		"bank_identifier_code",
		"salesforce_opportunity_id",
		"rssd_id",
		"dtc_sales_agent_part_num",
		"equity_ticker",
		"cmor_company_number",
		"lloyds_syndicate_performance",
		"cu_number",
		"equity_sedol",
		"ibm_number",
		"figi",
		"figi_previous",
		"figi_3",
		"figi_4",
		"figi_5",
		"moodys_deal_number",
		"legacy_deal_id",
	}
	basic := getBasicPrefix()
	return append(basic, prefix...)
}

func GetInstReferencePrefix() []string {
	prefix := []string {
		"cusip",
		"instrument_id",
		"mdy_debt_id",
		"moodys_legacy_id_number",
		"moodys_debt_id",
		"nomad_9char_issue_code",
		"org_id",
		"pid",
		"isin",
		"sedol",
		"bbgid",
		"bond_id",
		"common_code",
		"euroclear",
		"cedel",
		"nomad_alternate_issue_code",
		"loc_number",
		"common_code_3",
		"no_id_available",
		"aibd",
		"common_code_previous",
		"registration_number",
		"bloomberg_number",
		"sfg_tranche_number_dma",
		"hybrid",
		"insured_market",
		"ppn",
		"euro_fungible_at_emu",
		"enhanced_equipment_trust_certificates",
		"common_code_4",
		"salesforce_opportunity_id",
		"reference_bond_num",
		"reference_issuer_num",
		"cins",
		"seniority_change_date",
		"common_code_5",
		"project_finance",
		"additional_program_shelf_cusip",
		"program_shelf_cusip",
		"sec_registration_number",
		"st__st_prog_cusip",
		"sale_id",
		"vendor_cusip",
		"st_program_cusip",
		"lease_number",
		"rdb_debt_number",
		"deal_id",
		"orig_org_id",
	}
	basic := getBasicPrefix()
	return append(basic, prefix...)
}

func getBasicPrefix() []string {
	keys := []string{
		"cusip_6",
		"cusip_9",
		"mdy_id",
		"mdy_inst_id",
		"ma_id",
		"ma_inst_id",
		"mir_inst_id",
		"eq_isin",
		"global_isin",
		"eq_bbgid",
	}
	return keys
}  
