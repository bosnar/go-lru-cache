package lru

import (
	"container/list"
	"sync"
)

const CacheLimit = 100

// an interface for a CacheStore
type CacheStoreI interface {
	// Loader implements a function that load a value from the cache
	Loader(string) string
}

// CacheStore is LRU cache
type CacheStore struct {
	loader func(string) string
	cache  map[string]*list.Element
	queue  list.List
	mu     sync.Mutex
}

type entry struct {
	Key   string
	Value string
}

// New creates a new KeyStoreCache
func NewCacheStore(load CacheStoreI) *CacheStore {
	if load == nil {
		panic("load function is required")
	}
	return &CacheStore{
		loader: load.Loader,
		cache:  make(map[string]*list.Element),
	}
}

// Get the key from cache
func (k *CacheStore) Get(key string) string {
	k.mu.Lock()
	defer k.mu.Unlock()

	// Hit - move the item to the front of the queue
	if elem, hit := k.cache[key]; hit {
		k.queue.MoveToFront(elem)
		return elem.Value.(entry).Value
	}

	// Cache Miss - load and save it in cache
	p := entry{
		Key:   key,
		Value: k.loader(key),
	}

	// if cache is full remove the least used item
	if len(k.cache) >= CacheLimit {
		k.Del()
	}

	k.queue.PushFront(p)
	k.cache[key] = k.queue.Front()
	return p.Value
}

// Delete removes the least recently used item from the cache
func (k *CacheStore) Del() {
	elem := k.queue.Back()
	if elem != nil {
		k.queue.Remove(elem)
		delete(k.cache, elem.Value.(entry).Key)
	}
}
