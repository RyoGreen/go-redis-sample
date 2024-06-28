package server

import (
	"go-redis/controller"
	"net/http"
)

func Run() {
	entryHandler := controller.NewEntryController()
	userHandler := controller.NewUserController()
	jobHandler := controller.NewJobController()
	http.HandleFunc("/entries", entryHandler.ListEntries)
	http.HandleFunc("/users", userHandler.List)
	http.HandleFunc("/jobs", jobHandler.List)

	http.ListenAndServe(":8080", nil)
}
