package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"go-library-service/books"
	infr "go-library-service/books/infrastructure"
	"go-library-service/config"
	"go-library-service/helpers"
	"net/http"
	"os"
)

func main() {
	var settings config.Settings
	helpers.FromJson("config/config.json", &settings)
	service := books.NewService()
	router := mux.NewRouter()
	endpoints := books.MakeEndpoints(service)
	handler := infr.MakeHTTPHandler(router, endpoints)
	host := fmt.Sprintf("%s:%s", settings.HttpHost, settings.HttpPort)

	logger := log.NewLogfmtLogger(os.Stderr)
	logger.Log("err", http.ListenAndServe(host, handler))
}
