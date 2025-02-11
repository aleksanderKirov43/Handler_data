package transport

import (
	"handler-data/handler"
	"net/http"
)

func NewRouter() *http.ServeMux {

	router := http.NewServeMux()
	router.HandleFunc("/data", handler.GetDataHandler)
	return router
}
