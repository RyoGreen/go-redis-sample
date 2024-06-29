package server

import (
	"go-redis/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	m := mux.NewRouter()
	entryHandler := controller.NewEntryController()
	userHandler := controller.NewUserController()
	jobHandler := controller.NewJobController()

	m.HandleFunc("/users", userHandler.List).Methods(http.MethodGet)
	m.HandleFunc("/user/{id:[0-9]+}", userHandler.Get).Methods(http.MethodGet)
	m.HandleFunc("/user", userHandler.Create).Methods(http.MethodPost)
	m.HandleFunc("/user", userHandler.Update).Methods(http.MethodPut)
	m.HandleFunc("/user", userHandler.Delete).Methods(http.MethodDelete)

	m.HandleFunc("/jobs", jobHandler.List).Methods(http.MethodGet)
	m.HandleFunc("/job/{id:[0-9]+}", jobHandler.Get).Methods(http.MethodGet)
	m.HandleFunc("/job", jobHandler.Create).Methods(http.MethodPost)
	m.HandleFunc("/job", jobHandler.Update).Methods(http.MethodPut)
	m.HandleFunc("/job", jobHandler.Delete).Methods(http.MethodDelete)

	m.HandleFunc("/entries", entryHandler.List).Methods(http.MethodGet)
	m.HandleFunc("/entry/{id:[0-9]+}", entryHandler.Get).Methods(http.MethodGet)
	m.HandleFunc("/entry", entryHandler.Create).Methods(http.MethodPost)
	m.HandleFunc("/entry", entryHandler.Update).Methods(http.MethodPut)
	m.HandleFunc("/entry", entryHandler.Delete).Methods(http.MethodDelete)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Println(err)
	}
}
