package routes

import (
	"snakes_ladders/pkg/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router){
	router.HandleFunc("/addplayers", controllers.AddPlayer).Methods("POST")
	router.HandleFunc("/start", controllers.StartGame).Methods("GET")
	router.HandleFunc("/rolldice", controllers.DiceRoll).Methods("GET")
	router.HandleFunc("/positions", controllers.GetAllPlayersPositions).Methods("GET")
	router.HandleFunc("/position", controllers.GetPlayerPosition).Methods("GET")
	router.HandleFunc("/snl", controllers.SnlPositions).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
}
