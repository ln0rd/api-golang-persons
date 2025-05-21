package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ln0rd/api-golang-persons/models"
)

// Home controller
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(models.Personalities)
	if err != nil {
		http.Error(w, "Failed to encode personalities", http.StatusInternalServerError)
	}
}

func GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.ID) == id {
			err := json.NewEncoder(w).Encode(personality)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(models.Error{ErrorName: "Internal Server Error", Message: "Failed to encode personality", StatusCode: http.StatusInternalServerError})
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Error{ErrorName: "Not Found", Message: "Personality not found", StatusCode: http.StatusNotFound})
}
