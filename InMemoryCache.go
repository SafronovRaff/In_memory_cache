package In_memory_cache

import (
	"fmt"
	"time"
)

// In Memory Cache - a structure containing the duration of the cache storage and the cache itself
type InMemoryCashe struct {
	expireIn time.Duration
	cashe    map[string]CasheEntry
}

// Cache Entry - contains the time of setting the value and the value itselfе
type CasheEntry struct {
	settledAt time.Time
	value     interface{}
}

// NewCash - creates cache storage and cache storage time
func NewCash(expertIn time.Duration) *InMemoryCashe {
	return &InMemoryCashe{
		expireIn: expertIn,
		cashe:    make(map[string]CasheEntry),
	}
}

// Set - adds the value of saving by key along with the current time
func (c *InMemoryCashe) Set(key string, value interface{}) {
	c.cashe[key] = CasheEntry{
		settledAt: time.Now(),
		value:     value,
	}
}

// Get - returns the value by key
func (c *InMemoryCashe) Get(key string) interface{} {
	entry, ok := c.cashe[key]
	if ok && time.Since(entry.settledAt) <= c.expireIn {
		return entry.value
	}
	return fmt.Sprintf("The value has expired or it is not in the cache")
}

// Delete - delete the value by key
func (c *InMemoryCashe) Delete(key string) {
	_, ok := c.cashe[key]
	if ok {
		delete(c.cashe, key)
	}
}
