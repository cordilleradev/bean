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

type ConcurrentSet[T comparable] struct {
	mu sync.RWMutex
	s  map[T]struct{}
}

func NewConcurrentSet[T comparable]() *ConcurrentSet[T] {
	return &ConcurrentSet[T]{
		s: make(map[T]struct{}),
	}
}

func (c *ConcurrentSet[T]) Add(item T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.s[item] = struct{}{}
}

func (c *ConcurrentSet[T]) Remove(item T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.s, item)
}

func (c *ConcurrentSet[T]) Contains(item T) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.s[item]
	return exists
}

func (c *ConcurrentSet[T]) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.s)
}

func (c *ConcurrentSet[T]) Range(f func(item T) bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for k := range c.s {
		if !f(k) {
			break
		}
	}
}
