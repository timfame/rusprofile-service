package cache

import (
	"context"
	"encoding/json"
	"fmt"
	red "github.com/go-redis/redis/v8"
	"github.com/timfame/rusprofile-service/internal/config"
	"github.com/timfame/rusprofile-service/internal/models"
)

type redis struct {
	client *red.Client
	config *config.Cache
}

const (
	innKeyFmt = "inn:%s"
)

func NewRedis(ctx context.Context, redisCfg *config.Redis, cfg *config.Cache) (*redis, error) {
	client := red.NewClient(&red.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.DbNum,
	})

	_, err := client.Ping(ctx).Result()
	return &redis{
		client: client,
		config: cfg,
	}, err
}

func (r *redis) marshalSet(ctx context.Context, key string, value interface{}) error {
	mValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, mValue, r.config.ExpirationDuration).Err()
}

func (r *redis) unmarshalGet(ctx context.Context, key string, value interface{}) error {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(result), &value); err != nil {
		return err
	}
	return nil
}

func (r *redis) StoreCompany(ctx context.Context, company *models.Company) error {
	return r.marshalSet(ctx, fmt.Sprintf(innKeyFmt, company.INN), company)
}

func (r *redis) GetCompany(ctx context.Context, inn string) (*models.Company, error) {
	var company models.Company
	if err := r.unmarshalGet(ctx, fmt.Sprintf(innKeyFmt, inn), &company); err != nil {
		return nil, err
	}
	return &company, nil
}
