package grpc

import (
	rs "github.com/timfame/rusprofile-service/gen/rusprofile_service"
	"github.com/timfame/rusprofile-service/internal/models"
)

func companyToProto(c *models.Company) *rs.GetCompanyByINNResponse {
	if c == nil {
		return nil
	}
	return &rs.GetCompanyByINNResponse{
		Name:     c.Name,
		Inn:      c.INN,
		Kpp:      c.KPP,
		Director: c.Director,
	}
}
