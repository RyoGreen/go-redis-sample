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
	m.HandleFunc("/users/{id:[0-9]+}", userHandler.Get).Methods(http.MethodGet)
	m.HandleFunc("/users", userHandler.Create).Methods(http.MethodPost)
	m.HandleFunc("/users", userHandler.Update).Methods(http.MethodPut)
	m.HandleFunc("/users", userHandler.Delete).Methods(http.MethodDelete)

	m.HandleFunc("/jobs", jobHandler.List).Methods(http.MethodGet)
	m.HandleFunc("/job/{id:[0-9]+}", jobHandler.Get).Methods(http.MethodGet)
	m.HandleFunc("/jobs", jobHandler.Create).Methods(http.MethodPost)
	m.HandleFunc("/jobs", jobHandler.Update).Methods(http.MethodPut)
	m.HandleFunc("/jobs", jobHandler.Delete).Methods(http.MethodDelete)

	m.HandleFunc("/entries", entryHandler.ListEntries)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Println(err)
	}
}
