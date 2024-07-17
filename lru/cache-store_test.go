//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package lru

import (
	"strconv"
	"sync"
	"testing"
)


func TestCacheStore(t *testing.T) {

	loader := Loader{
		Data: GetMockDB(),
	}

	cache := NewCacheStore(&loader)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			value := cache.Get("Mock-" + strconv.Itoa(i))
			if value != "Mock-" + strconv.Itoa(i) {
				t.Errorf("failure %v", value)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	if len(cache.cache) != 100 {
		t.Errorf("length cache not 100: %d", len(cache.cache))
	}
	cache.Get("Mock-0")
	cache.Get("Mock-101")
	if _, ok := cache.cache["Mock-0"]; !ok {
		t.Errorf("delete incorrectly: %v", cache.cache)
	}

}
