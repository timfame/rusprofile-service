package grpc

import (
	"context"
	rs "github.com/timfame/rusprofile-service/gen/rusprofile_service"
)

func (s *server) GetCompanyByINN(ctx context.Context, req *rs.GetCompanyByINNRequest) (*rs.GetCompanyByINNResponse, error) {
	company, err := s.service.GetCompanyByINN(ctx, req.Inn)
	return companyToProto(company), err
}
