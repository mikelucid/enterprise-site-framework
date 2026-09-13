package cache

import (
	"sync"
	"time"
)

type Store interface {
	Get(key string) (string, bool)
	Set(key, value string, ttl time.Duration)
	Delete(key string)
}

type item struct {
	value     string
	expiresAt time.Time
}

type InMemory struct {
	mu   sync.RWMutex
	data map[string]item
}

func NewInMemory() *InMemory { return &InMemory{data: map[string]item{}} }

func (c *InMemory) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	it, ok := c.data[key]
	if !ok {
		return "", false
	}
	if !it.expiresAt.IsZero() && time.Now().After(it.expiresAt) {
		return "", false
	}
	return it.value, true
}

func (c *InMemory) Set(key, value string, ttl time.Duration) {
	exp := time.Time{}
	if ttl > 0 {
		exp = time.Now().Add(ttl)
	}
	c.mu.Lock()
	c.data[key] = item{value: value, expiresAt: exp}
	c.mu.Unlock()
}

func (c *InMemory) Delete(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}
