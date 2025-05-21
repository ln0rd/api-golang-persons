package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ln0rd/api-golang-persons/controllers"
)

func HandleRequest(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/personalities", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonalityByID).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
