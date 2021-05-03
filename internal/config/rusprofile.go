package config

import "github.com/timfame/rusprofile-service/pkg/env"

const (
	rusprofileBaseUrlEnv   = "RUSPROFILE_BASE_URL"
	rusprofileSearchUrlEnv = "RUSPROFILE_SEARCH_URL"

	rusprofileMainDivIDEnv                      = "RUSPROFILE_MAIN_DIV_ID"
	rusprofileSearchAmbiguousResultDivClassEnv  = "RUSPROFILE_SEARCH_AMBIGUOUS_RESULT_DIV_CLASS"
	rusprofileSearchEmptyResultDivClassEnv      = "RUSPROFILE_SEARCH_EMPTY_RESULT_DIV_CLASS"
	rusprofileSearchCompanyItemDivClassEnv      = "RUSPROFILE_SEARCH_COMPANY_ITEM_DIV_CLASS"
	rusprofileSearchCompanyItemTitleDivClassEnv = "RUSPROFILE_SEARCH_COMPANY_ITEM_TITLE_DIV_CLASS"

	rusprofileCompanyNameDivClassEnv                  = "RUSPROFILE_COMPANY_NAME_DIV_CLASS"
	rusprofileCompanyKPPSpanIDEnv                     = "RUSPROFILE_COMPANY_KPP_SPAN_ID"
	rusprofileCompanyInfoTitleClassEnv                = "RUSPROFILE_COMPANY_INFO_TITLE_CLASS"
	rusprofileCompanyInfoTitleDirectorTextEnv         = "RUSPROFILE_COMPANY_INFO_TITLE_DIRECTOR_TEXT"
	rusprofileCompanyInfoTextClassEnv                 = "RUSPROFILE_COMPANY_INFO_TEXT_CLASS"
)

type Rusprofile struct {
	BaseURL   string `json:"base_url"`
	SearchURL string `json:"search_url"`

	MainDivID                      string `json:"main_div_id"`
	SearchAmbiguousResultDivClass  string `json:"search_ambiguous_result_div_class"`
	SearchEmptyResultDivClass      string `json:"search_empty_result_div_class"`
	SearchCompanyItemDivClass      string `json:"search_company_item_div_class"`
	SearchCompanyItemTitleDivClass string `json:"search_company_item_title_div_class"`

	CompanyNameDivClass                  string `json:"company_name_div_class"`
	CompanyKPPSpanID                     string `json:"company_kpp_span_id"`
	CompanyInfoTitleClass                string `json:"company_info_title_class"`
	CompanyInfoTitleDirectorText         string `json:"company_info_title_director_text"`
	CompanyInfoTextClass                 string `json:"company_info_text_class"`
}

func (r *Rusprofile) Init() (err error) {
	r.BaseURL, err = env.GetString(rusprofileBaseUrlEnv)
	if err != nil {
		return
	}
	r.SearchURL, err = env.GetString(rusprofileSearchUrlEnv)
	if err != nil {
		return
	}
	r.MainDivID, err = env.GetString(rusprofileMainDivIDEnv)
	if err != nil {
		return
	}
	r.SearchAmbiguousResultDivClass, err = env.GetString(rusprofileSearchAmbiguousResultDivClassEnv)
	if err != nil {
		return
	}
	r.SearchEmptyResultDivClass, err = env.GetString(rusprofileSearchEmptyResultDivClassEnv)
	if err != nil {
		return
	}
	r.SearchCompanyItemDivClass, err = env.GetString(rusprofileSearchCompanyItemDivClassEnv)
	if err != nil {
		return
	}
	r.SearchCompanyItemTitleDivClass, err = env.GetString(rusprofileSearchCompanyItemTitleDivClassEnv)
	if err != nil {
		return
	}
	r.CompanyNameDivClass, err = env.GetString(rusprofileCompanyNameDivClassEnv)
	if err != nil {
		return
	}
	r.CompanyKPPSpanID, err = env.GetString(rusprofileCompanyKPPSpanIDEnv)
	if err != nil {
		return
	}
	r.CompanyInfoTitleClass, err = env.GetString(rusprofileCompanyInfoTitleClassEnv)
	if err != nil {
		return
	}
	r.CompanyInfoTitleDirectorText, err = env.GetString(rusprofileCompanyInfoTitleDirectorTextEnv)
	if err != nil {
		return err
	}
	r.CompanyInfoTextClass, err = env.GetString(rusprofileCompanyInfoTextClassEnv)
	return
}
