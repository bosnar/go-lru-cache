package main

import (
	"github.com/boss-ck/go-lru-cache/lru"
)

func main() {
	loader := lru.Loader{
		Data: lru.GetMockDB(),
	}

	cache := lru.NewCacheStore(&loader)

	_ = lru.RunMockData(cache)
}
