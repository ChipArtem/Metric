package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	usecase "github.com/ChipArtem/Metric/internal/server"
	"github.com/ChipArtem/Metric/internal/server/handlers"
	"github.com/ChipArtem/Metric/internal/server/repository"

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
	go func() {
		if err := http.ListenAndServe(host, mux); err != nil {
			log.Fatalf("start server: %v", err)
		}
	}()

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(
		signalChanel,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
LOOP:
	for {
		s := <-signalChanel
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Printf("\nRecive signal os: %v", s)
			break LOOP
		}
	}
}
