package server

import (
	"go-redis/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	m := mux.NewRouter()
	entryHandler := controller.NewEntryController()
	userHandler := controller.NewUserController()
	jobHandler := controller.NewJobController()

	m.HandleFunc("/users", userHandler.List).Methods(http.MethodGet)
	m.HandleFunc("/users/{id:[0-9]+}", userHandler.Get).Methods(http.MethodGet)
	m.HandleFunc("/users", userHandler.Create).Methods(http.MethodPost)
	m.HandleFunc("/users/{id:[0-9]+}", userHandler.Update).Methods(http.MethodPut)
	m.HandleFunc("/users", userHandler.Delete).Methods(http.MethodDelete)

	m.HandleFunc("/jobs", jobHandler.List).Methods(http.MethodGet)
	m.HandleFunc("/jobs/{id:[0-9]+}", jobHandler.Get).Methods(http.MethodGet)
	m.HandleFunc("/jobs", jobHandler.Create).Methods(http.MethodPost)
	m.HandleFunc("/jobs/{id:[0-9]+}", jobHandler.Update).Methods(http.MethodPut)
	m.HandleFunc("/jobs", jobHandler.Delete).Methods(http.MethodDelete)

	m.HandleFunc("/entries", entryHandler.ListEntries)

	http.ListenAndServe(":8080", m)
}
