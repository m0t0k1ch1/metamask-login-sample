package kvs

import "sync"

var (
	mu    sync.RWMutex           = sync.RWMutex{}
	cache map[string]interface{} = make(map[string]interface{})
)

func Set(key string, value interface{}) {
	mu.Lock()
	defer mu.Unlock()

	cache[key] = value
}

func Get(key string) (value interface{}, ok bool) {
	mu.RLock()
	defer mu.RUnlock()

	value, ok = cache[key]
	return
}
