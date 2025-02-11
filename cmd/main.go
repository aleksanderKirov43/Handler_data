package main

import (
	"handler-data/transport"
	"log"
	"net/http"
	"time"
)

// Вынести мапу на уровень репозитория. Раз в 5 сек нужно синковать кеш и базу данных, добавить TTL, использовать внутренний кеш,
//все функции должны быть доступны на уровне меён пакета. handler/service/repository

func main() {

	router := transport.NewRouter()

	go func() {
		for {
			time.Sleep(5 * time.Second)
		}
	}()

	err := http.ListenAndServe(":7777", router)
	if err != nil {
		log.Fatal(err)
	}
}
