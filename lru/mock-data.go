package lru

import (
	"strconv"
	"sync"
	"sync/atomic"
)

type MockData struct {
	Counter int32
}

// use a mock data to simulate a database
func (db *MockData) Get(key string) string {
	atomic.AddInt32(&db.Counter, 1)
	return key
}

func GetMockDB() *MockData {
	return &MockData{
		Counter: 0,
	}
}

func RunMockData(cache *CacheStore) error {

	wg := new(sync.WaitGroup)

	// run worker
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = cache.Get(`Mock-` + strconv.Itoa(i))
		}()
	}

	wg.Wait()

	return nil
}
