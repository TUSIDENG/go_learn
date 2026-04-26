package sync_example

import (
	"sync"
)

type SafeCounter struct {
	mu     sync.RWMutex
	counts map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		counts: make(map[string]int),
	}
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counts[key]
}

func (c *SafeCounter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts = make(map[string]int)
}

func (c *SafeCounter) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	keys := make([]string, 0, len(c.counts))
	for k := range c.counts {
		keys = append(keys, k)
	}
	return keys
}

func (c *SafeCounter) BatchIncrement(key string, total int, workers int) {
	if workers <= 0 {
		workers = 1
	}

	var wg sync.WaitGroup
	wg.Add(workers)

	chunk := total / workers
	remainder := total % workers

	for i := 0; i < workers; i++ {
		adds := chunk
		if i == 0 {
			adds += remainder
		}

		go func(n int) {
			defer wg.Done()
			for j := 0; j < n; j++ {
				c.Inc(key)
			}
		}(adds)
	}

	wg.Wait()
}
