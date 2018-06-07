package safemap

import "sync"
import "fmt"
import "github.com/gogather/com"

// SafeMap safe map struct
type SafeMap struct {
	sync.RWMutex
	M map[string]interface{} `json:"_"`
}

// New new a SafeMap
func New() *SafeMap {
	return &SafeMap{
		M: make(map[string]interface{}),
	}
}

// Put put element into safemap
func (sm *SafeMap) Put(key string, value interface{}) {
	sm.Lock()
	sm.M[key] = value
	sm.Unlock()
}

// Remove remove element from safemap
func (sm *SafeMap) Remove(key string) {
	sm.Lock()
	delete(sm.M, key)
	sm.Unlock()
}

// Get get element from safemap
func (sm *SafeMap) Get(key string) (interface{}, bool) {
	defer func() {
		sm.RUnlock()
	}()
	sm.RLock()
	v, ok := sm.M[key]
	return v, ok
}

func (sm *SafeMap) String() string {
	defer func() {
		sm.RUnlock()
	}()
	sm.RLock()
	return fmt.Sprintf("%v", sm.M)
}

// JSON convert map to json string
func (sm *SafeMap) JSON() (json string) {
	defer func() {
		sm.RUnlock()
	}()
	sm.RLock()
	json, _ = com.JsonEncode(sm.M)
	return
}

// GetMap get original map
func (sm *SafeMap) GetMap() map[string]interface{} {
	return sm.M
}

// Clear clear the map
func (sm *SafeMap) Clear() {
	sm.M = make(map[string]interface{})
}
