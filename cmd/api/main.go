package main

import (
	"fmt"
	"net/http"

	"github.com/Sn0wye/go-learn/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	
	fmt.Println("Server is running")

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Error(err)
	}
}

