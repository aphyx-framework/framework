package cache

import (
	"github.com/muesli/cache2go"
	"github.com/rama-adi/RyFT-Framework/app"
	"time"
)

type Table struct {
	Auth *cache2go.CacheTable
}

func Init() Table {
	return Table{
		Auth: cache2go.Cache("auth"),
	}
}

// CacheOrMake will cache the data if it's not cached yet, or return the cached data if it's already cached
func (_ Table) CacheOrMake(
	table *cache2go.CacheTable,
	key string,
	f func() (interface{}, error, time.Duration),
) (interface{}, error) {
	if app.Config.Caching.Enabled == false {
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
