package repository

import (
	"sync"
	"time"
)

type Cache struct {
	mapData           map[string]string
	mutex             sync.Mutex
	defaultExpiration time.Duration // Продолжительность жизнит кеша
	cleanupInterval   time.Duration // Интервал, через который кеш будет очищен
}

func NewCache(expiriation, interval time.Duration) *Cache {
	return &Cache{
		data:              make(map[string]string),
		defaultExpiration: expiriation,
		cleanupInterval:   interval,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	val, found := c.data[key]
	return val, found
}

func Set(key string, ) {

}
