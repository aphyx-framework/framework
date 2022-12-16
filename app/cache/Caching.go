package cache

import (
	"github.com/muesli/cache2go"
	"github.com/rama-adi/RyFT-Framework/framework/caching"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
)

type Table struct {
	Auth   *cache2go.CacheTable
	config configuration.Configuration
}

type UserTable struct {
	Auth caching.CacheTable
}

func InitializeCacheTable(config configuration.Configuration) UserTable {
	return UserTable{
		Auth: caching.CacheTable{
			Table:  cache2go.Cache("auth"),
			Config: config,
		},
	}
}
