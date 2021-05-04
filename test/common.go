package test

import (
	"github.com/timfame/rusprofile-service/internal/cache"
	"github.com/timfame/rusprofile-service/internal/config"
	"github.com/timfame/rusprofile-service/internal/models"
	"github.com/timfame/rusprofile-service/internal/rusprofile"
	"github.com/timfame/rusprofile-service/pkg/logger"
	"testing"
)

func checkCompany(t *testing.T, expected, actual *models.Company) {
	if expected.Name != actual.Name {
		t.Fatalf("Wrong company name:\n\texpected: %s,\n\tfound: %s", expected.Name, actual.Name)
	}
	if expected.INN != actual.INN {
		t.Fatalf("Wrong company INN:\n\texpected: %s,\n\tfound: %s", expected.INN, actual.INN)
	}
	if expected.KPP != actual.KPP {
		t.Fatalf("Wrong company KPP:\n\texpected: %s,\n\tfound: %s", expected.KPP, actual.KPP)
	}
	if expected.Director != actual.Director {
		t.Fatalf("Wrong company director:\n\texpected: %s,\n\tfound: %s", expected.Director, actual.Director)
	}
}

func getTestServiceAndLogger() (rusprofile.Service, *logger.Logger) {
	rusprofileCfg := &config.Rusprofile{
		BaseURL:                        "https://www.rusprofile.ru",
		SearchURL:                      "https://www.rusprofile.ru/search?query=",
		MainDivID:                      "main",
		SearchAmbiguousResultDivClass:  "company-main search-result__main",
		SearchEmptyResultDivClass:      "main-content search-result emptyresult",
		SearchCompanyItemDivClass:      "company-item",
		SearchCompanyItemTitleDivClass: "company-item__title",
		CompanyNameDivClass:            "company-name",
		CompanyKPPSpanID:               "clip_kpp",
		CompanyInfoTitleClass:          "company-info__title",
		CompanyInfoTitleDirectorText:   "Руководитель",
		CompanyInfoTextClass:           "company-info__text",
	}
	l := logger.New(logger.WithDebugLevel(), logger.WithServiceName("rusprofile-test"))
	return rusprofile.NewLogger(rusprofile.NewBase(rusprofileCfg, cache.NewHashmap(&config.Cache{
		ExpirationDuration: 0,
		ClearingInterval:   0,
	})), l), l
}
