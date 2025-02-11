package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"handler-data/repository"
)

var cache = repository.NewSafeCache()

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
		cache.Set("data", data, 10*time.Minute)
		value = data
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}
