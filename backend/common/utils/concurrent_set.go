package utils

import "sync"

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
