package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// Вынести мапу на уровень репозитория. Раз в 5 сек нужно синковать кеш и базу данных, добавить TTL, использовать внутренний кеш,
//все функции должны быть доступны на уровне меён пакета. handler/service/repository

func main() {
	m := http.NewServeMux()

	var mutex sync.RWMutex

	limiter := time.Tick(1 * time.Second)

	m.HandleFunc("/data", func(writer http.ResponseWriter, request *http.Request) {

		<-limiter

		mutex.Lock()
		defer mutex.Unlock()

		writer.Header().Set("Content-Type", "application/json")

		mapData := make(map[string]string)

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
