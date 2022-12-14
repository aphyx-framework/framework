package cache

import (
	"github.com/muesli/cache2go"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"time"
)

type Table struct {
	Auth   *cache2go.CacheTable
	config configuration.Configuration
}

func Init(config configuration.Configuration) Table {
	return Table{
		Auth:   cache2go.Cache("auth"),
		config: config,
	}
}

// CacheOrMake will cache the data if it's not cached yet, or return the cached data if it's already cached
func (t Table) CacheOrMake(
	table *cache2go.CacheTable,
	key string,
	f func() (interface{}, error, time.Duration),
) (interface{}, error) {
	if t.config.Caching.Enabled == false {
		val, err, _ := f()
		return val, err
	}

	// Simple caching function
	if table.Exists(key) {
		cacheValue, cacheError := table.Value(key)
		return cacheValue.Data(), cacheError
	} else {
		data, err, expiry := f()

		if err != nil {
			return nil, err
		}

		table.Add(key, expiry, data)
		return data, nil
	}

}

func (t Table) BustCache(
	table *cache2go.CacheTable,
	key string,
) error {

	if t.config.Caching.Enabled == false {
		return nil
	}

	_, err := table.Delete(key)
	return err
}
