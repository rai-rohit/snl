package main

import (
	"fmt"
	"log"
	"net/http"
	"snakes_ladders/pkg/config"
	"snakes_ladders/pkg/database"
	"snakes_ladders/pkg/routes"
	"snakes_ladders/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	config, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	
	err = database.Connect(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DB connected successfully!!\n")

	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9090", middleware.Authorize(router)))
}
