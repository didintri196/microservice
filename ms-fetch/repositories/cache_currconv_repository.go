package repositories

import (
	"errors"
	"ms-fetch/domain/constants"
	"ms-fetch/domain/interfaces"
	"ms-fetch/domain/models"

	"github.com/patrickmn/go-cache"
)

type CacheCurrConvRepository struct {
	Cache *cache.Cache
}

func NewCacheCurrConvRepo(cache *cache.Cache) interfaces.ICacheCurrConvRepository {
	return &CacheCurrConvRepository{
		Cache: cache,
	}
}

func (repo CacheCurrConvRepository) GetIDRtoUSD() (model models.Currency, err error) {
	// Get the string associated with the key "constants.KeyIDRtoUSD" from the cache
	dataCache, found := repo.Cache.Get(constants.KeyIDRtoUSD)
	if found {
		model.IdrUsd = dataCache.(float64)
		return model, nil
	}
	return model, errors.New("key not found")
}

func (repo CacheCurrConvRepository) SetIDRtoUSD(currency models.Currency) {
	repo.Cache.Set(constants.KeyIDRtoUSD, currency.IdrUsd, cache.DefaultExpiration)
}
