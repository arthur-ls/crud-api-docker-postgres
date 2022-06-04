package route

import (
	"github.com/arthur-ls/project1/controller"
	"github.com/gorilla/mux"
)

func Route(router *mux.Router) {
	router.HandleFunc("/create", controller.HandleInsertByID).Methods("POST", "OPTIONS")
	router.HandleFunc("/get", controller.HandleGetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/get/{id}", controller.HandleGetID).Methods("GET", "OPTIONS")
	router.HandleFunc("/update/{id}", controller.HandleUpdateID).Methods("PUT", "OPTIONS")
	router.HandleFunc("/delete/{id}", controller.HandleDeleteID).Methods("DELETE", "OPTIONS")
}
