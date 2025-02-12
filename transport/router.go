package transport

import (
	"handler-data/repository"

	"net/http"
)

func NewRouter() *http.ServeMux {

	router := http.NewServeMux()
	router.HandleFunc("/data", repository.GetDataHandler)
	return router
}
