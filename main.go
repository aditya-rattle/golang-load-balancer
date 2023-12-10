package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"loadbalancer.com/controller"
	"loadbalancer.com/models"
	"loadbalancer.com/service"
)

func main() {

	router := mux.NewRouter()

	allAvailableServers := []*models.Server{
		models.NewServer("https://www.google.com"),
		models.NewServer("https://www.facebook.com"),
		models.NewServer("https://www.duckduckgo.com"),
	}

	lb := service.NewLoadBalancer(allAvailableServers)
	ctr := controller.NewController(lb)

	router.HandleFunc("/", ctr.RouteRequest).Methods("GET")

	log.Println("Server starting at 8080")

	http.ListenAndServe(":8080", router)
}
