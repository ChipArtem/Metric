package main

import (
	"log"
	"net/http"

	"MyMetric/internal/server/handlers"
	"MyMetric/internal/server/repository"
	"MyMetric/internal/server/usecase"

	"github.com/gorilla/mux"
)

func main() {

	host := "127.0.0.1:8080"
	repo := repository.NewRepoMem()
	bl := usecase.NewMetricBusinessLogic(repo)
	handlers := handlers.NewMetricHandler(bl, host)

	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.GetAll).Methods("GET")
	mux.HandleFunc("/update/{mtype}/{name}/{value}", handlers.SetMetric).Methods("POST")
	mux.HandleFunc("/value/{mtype}/{name}", handlers.GetMetric).Methods("GET")
	mux.Use(handlers.MiddlewareCheckHost)
	if err := http.ListenAndServe(host, mux); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
