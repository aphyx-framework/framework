package caching

import (
	"github.com/muesli/cache2go"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"time"
)

type CacheTable struct {
	Table  *cache2go.CacheTable
	Config configuration.Configuration
}

func (cacheTable CacheTable) CacheOrMake(
	key string,
	f func() (interface{}, error, time.Duration),
) (interface{}, error) {
	if cacheTable.Config.Caching.Enabled == false {
		val, err, _ := f()
		return val, err
	}

	// Simple caching function
	if cacheTable.Table.Exists(key) {
		cacheValue, cacheError := cacheTable.Table.Value(key)
		return cacheValue.Data(), cacheError
	} else {
		data, err, expiry := f()

		if err != nil {
			return nil, err
		}

		cacheTable.Table.Add(key, expiry, data)
		return data, nil
	}
}

func (cacheTable CacheTable) BustCache(key string) error {
	if cacheTable.Config.Caching.Enabled == false {
		return nil
	}
	_, err := cacheTable.Table.Delete(key)
	return err
}
