package test

import (
	"context"
	"github.com/timfame/rusprofile-service/internal/models"
	"github.com/timfame/rusprofile-service/internal/rusprofile"
	"testing"
)

func TestUniqueSearchResult(t *testing.T) {
	service, _ := getTestServiceAndLogger()
	company, err := service.GetCompanyByINN(context.Background(), "7813045547")
	if err != nil {
		t.Fatal(err)
	}
	checkCompany(t, &models.Company{
		Name:     "ФЕДЕРАЛЬНОЕ ГОСУДАРСТВЕННОЕ АВТОНОМНОЕ ОБРАЗОВАТЕЛЬНОЕ УЧРЕЖДЕНИЕ ВЫСШЕГО ОБРАЗОВАНИЯ \"НАЦИОНАЛЬНЫЙ ИССЛЕДОВАТЕЛЬСКИЙ УНИВЕРСИТЕТ ИТМО\"",
		Director: "Васильев Владимир Николаевич",
		INN:      "7813045547",
		KPP:      "781301001",
	}, company)
}

func TestMultipleSearchResult(t *testing.T) {
	service, _ := getTestServiceAndLogger()

	company, err := service.GetCompanyByINN(context.Background(), "7843007274")
	if err != nil {
		t.Fatal(err)
	}
	checkCompany(t, &models.Company{
		Name:     "ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ \"БУЛОЧНАЯ  № 1\"",
		Director: "Цапиков Алексей Вадимович",
		INN:      "7843007274",
		KPP:      "783901001",
	}, company)

	company, err = service.GetCompanyByINN(context.Background(), "7707083893")
	if err != nil {
		t.Fatal(err)
	}
	checkCompany(t, &models.Company{
		Name:     "ПУБЛИЧНОЕ АКЦИОНЕРНОЕ ОБЩЕСТВО \"СБЕРБАНК РОССИИ\"",
		Director: "Греф Герман Оскарович",
		INN:      "7707083893",
		KPP:      "773601001",
	}, company)
}

func TestEmptySearchResult(t *testing.T) {
	service, _ := getTestServiceAndLogger()
	_, err := service.GetCompanyByINN(context.Background(), "1421412313")
	if err != rusprofile.ErrNotFound {
		t.Fatalf("Returned error must be ErrNotFound, but found: %v", err)
	}
}
