package cache

import (
	"github.com/muesli/cache2go"
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

func (_ Table) CacheOrMake(
	table *cache2go.CacheTable,
	key string,
	f func() (interface{}, error, time.Duration),
) (interface{}, error) {
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
