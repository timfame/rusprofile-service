package cache

import (
	"context"
	"github.com/pkg/errors"
	"github.com/timfame/rusprofile-service/internal/models"
)

type Storage interface {
	StoreCompany(ctx context.Context, company *models.Company) error
	GetCompany(ctx context.Context, inn string) (*models.Company, error)
}

var (
	ErrNotFound = errors.New("cannot found in cache items")
)
