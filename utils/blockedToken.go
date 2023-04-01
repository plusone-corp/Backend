package utils

import (
	"sync"
	"time"
)

type LoggedOutToken struct {
	items map[string]cacheItem
	mutex sync.RWMutex
}

type cacheItem struct {
	token      string
	expiration time.Time
}

func NewBlockToken() *LoggedOutToken {
	c := &LoggedOutToken{
		items: make(map[string]cacheItem),
	}
	go c.cleanup()
	return c
}

func (c *LoggedOutToken) Set(key, value string, expiration time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items[key] = cacheItem{
		token:      value,
		expiration: expiration,
	}
}

func (c *LoggedOutToken) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.items[key]
	if !found {
		return "", false
	}
	if time.Now().After(item.expiration) {
		delete(c.items, key)
		return "", false
	}
	return item.token, true
}

func (c *LoggedOutToken) cleanup() {
	for {
		<-time.After(30 * time.Minute)

		c.mutex.Lock()
		for key, item := range c.items {
			if time.Now().After(item.expiration) {
				delete(c.items, key)
			}
		}
		c.mutex.Unlock()
	}
}
