package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rishavpreet/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service")

	fmt.Println(`

	`)

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
