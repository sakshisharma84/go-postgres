package router

import (
    "go-postgres/handlers"
    "github.com/gorilla/mux"
)

func Router() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/api/vehicle/{id}", handlers.GetVehicle).Methods("GET")
    router.HandleFunc("/api/vehicle", handlers.GetAllVehicle).Methods("GET")
    router.HandleFunc("/api/newvehicle", handlers.CreateVehicle).Methods("POST")
    router.HandleFunc("/api/vehicle/{id}", handlers.UpdateVehicle).Methods("PUT")
    router.HandleFunc("/api/deletevehicle/{id}", handlers.DeleteVehicle).Methods("DELETE")
    router.HandleFunc("/api/search/{query}", handlers.SearchVehicle).Methods("GET")

    return router
}
