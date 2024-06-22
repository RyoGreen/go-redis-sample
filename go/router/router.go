package router

import (
	"go-redis/controller"
	"net/http"
)

func Run() {
	entryHandler := controller.NewEntryController()

	http.HandleFunc("/entries", entryHandler.ListEntries)
}
