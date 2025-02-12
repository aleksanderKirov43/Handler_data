package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Значения, для хранения в кеше
type CacheEntry struct {
	value      map[string]string
	expiration int64
}

// Кэш
type SafeCache struct {
	syncMap sync.Map
}

var cache = NewSafeCache()

func GetDataHandler(w http.ResponseWriter, r *http.Request) {

	value, found := cache.Get("data")
	if !found {
		data := map[string]string{
			"google":     "google.com",
			"yahoo!":     "search.yahoo.com",
			"yandex":     "yandex.com",
			"duckduckgo": "duckduckgo.com",
			"baidu":      "baidu.com",
			"bing":       "bing.com",
			"ask":        "ask.com",
			"archive":    "archive.org",
			"ecosia":     "ecosia.org",
		}
		cache.Set("data", data, 1*time.Minute)
		value = data

		fmt.Println("Кеш валидирован")
	} else {
		fmt.Println("Кеш не валидирован")
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func NewSafeCache() *SafeCache {
	return &SafeCache{}
}

// Сохраняем значения с заданными TTL
// Задаём "врмя жизни" в секундах.
func (sc *SafeCache) Set(key string, data map[string]string, ttl time.Duration) {
	expiration := time.Now().Add(ttl).UnixNano()
	sc.syncMap.Store(key, CacheEntry{value: data, expiration: expiration})

}

// Извлекаем значение из кэша, если оно не найдено
// или истёк срок действия.
func (sc *SafeCache) Get(key string) (map[string]string, bool) {
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
