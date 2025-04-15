package main

import (
	"fmt"
	"handler-data/transport"

	"log"
	"net/http"
)

func main() {

	router := transport.NewRouter()

	fmt.Println("HTTP Сервер прослушивает порт : '7777'")

	err := http.ListenAndServe(":7777", router)
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера : %v", err)
	}
}
