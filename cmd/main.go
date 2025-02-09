package main

import (
	"encoding/json"
	"sync"
	"time"

	"net/http"
)

func main() {
	m := http.NewServeMux()

	var mutex sync.Mutex

	limiter := time.Tick(1 * time.Second)

	m.HandleFunc("/data", func(writer http.ResponseWriter, request *http.Request) {

		<-limiter

		if request.Method != http.MethodGet {
			http.Error(writer, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		writer.Header().Set("Content-Type", "application/json")
		var mapData map[string]string

		mapData = map[string]string{
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

		jsonResponse, err := json.Marshal(mapData)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Write(jsonResponse)
	})

	err := http.ListenAndServe(":7777", m)
	if err != nil {
		return
	}
}
