package kvs

import "sync"

var (
	mu    sync.RWMutex           = sync.RWMutex{}
	store map[string]interface{} = make(map[string]interface{})
)

func Set(key string, value interface{}) {
	mu.Lock()
	defer mu.Unlock()

	store[key] = value
}

func Get(key string) (value interface{}, ok bool) {
	mu.RLock()
	defer mu.RUnlock()

	value, ok = store[key]
	return
}

func Delete(key string) (value interface{}, ok bool) {
	mu.Lock()
	defer mu.Unlock()

	value, ok = store[key]
	delete(store, key)
	return
}
