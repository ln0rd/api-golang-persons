package routes

import (
	"github.com/gorilla/mux"
	"github.com/ln0rd/api-golang-persons/controllers"
	"github.com/ln0rd/api-golang-persons/middleware"
	"go.uber.org/zap"
)

type Route struct {
	logger *zap.Logger
	pc     *controllers.PersonalityController
}

func NewRoute(logger *zap.Logger, pc *controllers.PersonalityController) *Route {
	return &Route{logger: logger, pc: pc}
}

func (r *Route) SetupRoutes(router *mux.Router) {
	// Aplicando o middleware
	router.Use(middleware.SetHeaders)

	router.HandleFunc("/", r.pc.Home).Methods("GET")
	router.HandleFunc("/personalities", r.pc.GetAllPersonalities).Methods("GET")
	router.HandleFunc("/personalities/{id}", r.pc.GetPersonalityByID).Methods("GET")
	router.HandleFunc("/personalities", r.pc.CreatePersonality).Methods("POST")
	router.HandleFunc("/personalities/{id}", r.pc.DeletePersonality).Methods("DELETE")
	router.HandleFunc("/personalities/{id}", r.pc.UpdatePersonality).Methods("PUT")
}
