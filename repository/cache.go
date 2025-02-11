package repository

import (
	"sync"
	"time"
)

// Значения, для хранения в кеше
type CacheEntry struct {
	value      interface{}
	expiration int64
}

// Кэш
type SafeCache struct {
	syncMap sync.Map
}

func NewSafeCache() *SafeCache {
	return &SafeCache{}
}

// Сохраняем значения с заданными TTL
// Задаём "врмя жизни" в секундах.
func (sc *SafeCache) Set(key string, value interface{}, ttl time.Duration) {
	expiration := time.Now().Add(ttl).UnixNano()
	sc.syncMap.Store(key, CacheEntry{value: value, expiration: expiration})
}

// Извлекаем значение из кэша, если оно не найдено
// или истёк срок действия.
func (sc *SafeCache) Get(key string) (interface{}, bool) {
	entry, found := sc.syncMap.Load(key)
	if !found {
		return nil, false
	}

	cacheEntry := entry.(CacheEntry)
	if time.Now().UnixNano() > cacheEntry.expiration {
		sc.syncMap.Delete(key)
		return nil, false
	}
	return cacheEntry.value, true
}

// Удаление значения из кэша
func (sc *SafeCache) Delete(key string) {
	sc.syncMap.Delete(key)
}

// Переодическая очистка кэша
func (sc *SafeCache) CleanUp() {
	for {
		time.Sleep(1 * time.Minute)
		sc.syncMap.Range(func(key, entry interface{}) bool {
			cacheEntry := entry.(CacheEntry)
			if time.Now().UnixNano() > cacheEntry.expiration {
				sc.syncMap.Delete(key)
			}
			return true
		})
	}
}
