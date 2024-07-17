package lru

type Loader struct {
	Data *MockData
}

// implements a function that loads a value from the cache
func (l *Loader) Loader(key string) string {
	return l.Data.Get(key)
}
