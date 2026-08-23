package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.data[key]

	return value, ok
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

func main() {

	cache := NewCache()

	const (
		readerCount = 5
		writerCount = 3
		deleteCount = 1
	)

	var wg sync.WaitGroup
	wg.Add(readerCount + writerCount + deleteCount)

	for writerID := range writerCount {
		go func(id int) {
			defer wg.Done()
			for j := range 100 {
				key := fmt.Sprintf("worker-%d", id)
				value := fmt.Sprintf("value-%d", j)
				cache.Set(key, value)
			}
		}(writerID)
	}

	for range readerCount {
		go func() {
			defer wg.Done()
			for range 100 {
				cache.Get("worker-1")
				cache.Get("worker-2")
			}
		}()
	}

	for range deleteCount {
		go func() {
			defer wg.Done()
			cache.Delete("worker-1")
		}()
	}

	wg.Wait()

	value, ok := cache.Get("worker-1")
	fmt.Println("worker-1:", value, ok)

	value, ok = cache.Get("worker-2")
	fmt.Println("worker-2:", value, ok)
}
