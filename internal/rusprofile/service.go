package rusprofile

import (
	"context"
	"errors"
	"github.com/timfame/rusprofile-service/internal/models"
)

type Service interface {
	GetCompanyByINN(ctx context.Context, inn string) (*models.Company, error)
}

var (
	ErrNotFound                = errors.New("cannot find a company with this INN")
	ErrInternalRusprofile      = errors.New("internal rusprofile.ru error")
	ErrTooManyRequests         = errors.New("too many requests to rusprofile.ru")
	ErrRuspofileResponseStatus = errors.New("rusprofile.ru response status code is not OK")
)
