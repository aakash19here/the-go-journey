package main

import (
	"fmt"
	"sync"
)

type ConfigCache struct {
	data     map[string]string
	dataLock sync.RWMutex
}

func (c *ConfigCache) Set(key, value string) {
	c.dataLock.Lock()
	defer c.dataLock.Unlock()
	c.data[key] = value
}

func (c *ConfigCache) Get(key string) (string, bool) {
	c.dataLock.RLock()
	defer c.dataLock.RUnlock()
	val, exists := c.data[key]
	return val, exists
}

func main() {
	cache := &ConfigCache{
		data: make(map[string]string),
	}

	var wg sync.WaitGroup

	for i := range 100 {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			for range 50 {
				cache.Get("database_url")
			}
		}(i)
	}

	for i := range 10 {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for range 50 {
				cache.Set("database_url", fmt.Sprintf("postgres://db-%d", id))
			}
		}(i)
	}

	wg.Wait()

	finalVal, _ := cache.Get("database_url")
	fmt.Println("Operations complete. Final database URL:", finalVal)

}
