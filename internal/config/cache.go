package config

import (
	"github.com/timfame/rusprofile-service/pkg/env"
	"time"
)

const (
	cacheExpirationDurationEnv = "CACHE_EXPIRATION_DURATION_SECONDS"
	cacheClearingIntervalEnv   = "CACHE_CLEARING_INTERVAL_SECONDS"
)

type Cache struct {
	ExpirationDuration time.Duration `json:"expiration_duration_seconds"`
	ClearingInterval   time.Duration `json:"clearing_interval_seconds"`
}

func (c *Cache) Init() (err error) {
	c.ExpirationDuration, err = env.GetSecondsDuration(cacheExpirationDurationEnv)
	if err != nil {
		return
	}
	c.ClearingInterval, err = env.GetSecondsDuration(cacheClearingIntervalEnv)
	return
}
