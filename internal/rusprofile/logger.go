package rusprofile

import (
	"context"
	"github.com/timfame/rusprofile-service/internal/models"
	log "github.com/timfame/rusprofile-service/pkg/logger"
	"go.uber.org/zap"
)

type logger struct {
	service Service
	logger  *log.Logger
}

func NewLogger(s Service, l *log.Logger) *logger {
	return &logger{
		service: s,
		logger:  l,
	}
}

func (l *logger) GetCompanyByINN(ctx context.Context, inn string) (company *models.Company, err error) {
	l.logger.Info("Request GetCompanyByINN", zap.String("inn", inn))
	company, err = l.service.GetCompanyByINN(ctx, inn)
	if err != nil {
		l.logger.Error("Response GetCompanyByINN failed", zap.Error(err))
	} else {
		l.logger.Info("Response GetCompanyByINN",
			zap.String("name", company.Name),
			zap.String("inn", company.INN),
			zap.String("kpp", company.KPP),
			zap.String("director", company.Director))
	}
	return
}
