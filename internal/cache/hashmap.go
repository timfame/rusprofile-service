package cache

import (
	"context"
	"github.com/timfame/rusprofile-service/internal/config"
	"github.com/timfame/rusprofile-service/internal/models"
	"sync"
	"time"
)

type hashmap struct {
	items  map[string]item
	mu     sync.RWMutex
	config *config.Cache
}

type item struct {
	value    interface{}
	expireAt int64
}

func (i *item) isExpired() bool {
	return time.Now().Unix() > i.expireAt
}

func NewHashmap(cfg *config.Cache) *hashmap {
	hm := &hashmap{
		items:  make(map[string]item),
		mu:     sync.RWMutex{},
		config: cfg,
	}
	if cfg.ExpirationDuration > 0 && cfg.ClearingInterval > 0 {
		go hm.startClearer()
	}
	return hm
}

func (h *hashmap) startClearer() {
	for {
		<-time.After(h.config.ClearingInterval)
		h.mu.Lock()
		for k, v := range h.items {
			if v.isExpired() {
				delete(h.items, k)
			}
		}
		h.mu.Unlock()
	}
}

func (h *hashmap) set(key string, value interface{}) error {
	item := item{
		value:    value,
		expireAt: time.Now().Add(h.config.ExpirationDuration).Unix(),
	}
	h.mu.Lock()
	h.items[key] = item
	h.mu.Unlock()
	return nil
}

func (h *hashmap) get(key string) (item, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if i, ok := h.items[key]; ok {
		if h.config.ExpirationDuration > 0 && i.isExpired() {
			return item{}, ErrNotFound
		}
		return i, nil
	}
	return item{}, ErrNotFound
}

func (h *hashmap) StoreCompany(ctx context.Context, company *models.Company) error {
	return h.set(company.INN, company)
}

func (h *hashmap) GetCompany(ctx context.Context, inn string) (*models.Company, error) {
	item, err := h.get(inn)
	if err != nil {
		return nil, err
	}
	if company, ok := item.value.(*models.Company); ok {
		return company, nil
	}
	return nil, ErrNotFound
}
