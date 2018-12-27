package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/brianvoe/gofakeit"
)

func main() {
	log.Println("Starting random paragraph service")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	r := mux.NewRouter()
	r.HandleFunc("/ready", ReadinessHandler).Methods("GET")
	r.HandleFunc("/liveness", LivenessHandler).Methods("GET")
	r.HandleFunc("/random", RandomHandler).Methods("GET")

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	log.Println("Ready to serve")

	shutdown := make(chan struct{}, 1)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			shutdown <- struct{}{}
			log.Printf("%v", err)
		}
	}()
	log.Print("The service is ready to listen and serve.")

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}
	case <-shutdown:
		log.Printf("Got an error...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	data := gofakeit.Paragraph(1, 3, 10, " ")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))

}

func LivenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
