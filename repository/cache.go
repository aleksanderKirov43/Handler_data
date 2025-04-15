package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	value, err := rdb.Get(ctx, "data").Result()
	if err == redis.Nil {
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
		jsonData, _ := json.Marshal(data)
		err := rdb.Set(ctx, "data", jsonData, 10*time.Minute).Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		value = string(jsonData)
		fmt.Println("Кеш валидирован")
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Кеш не валидирован")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(value))
}
