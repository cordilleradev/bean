package utils

import (
	"sync"
)

type ConcurrentMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		m: make(map[K]V),
	}
}

func (c *ConcurrentMap[K, V]) Store(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *ConcurrentMap[K, V]) Load(key K) (value V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok = c.m[key]
	return
}

func (c *ConcurrentMap[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}

func (c *ConcurrentMap[K, V]) Range(f func(key K, value V) bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for k, v := range c.m {
		if !f(k, v) {
			break
		}
	}
}

func (c *ConcurrentMap[K, V]) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.m)
}
