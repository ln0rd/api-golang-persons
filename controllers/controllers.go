package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ln0rd/api-golang-persons/errs"
	"github.com/ln0rd/api-golang-persons/models"
	"github.com/ln0rd/api-golang-persons/services"
	"go.uber.org/zap"
)

type PersonalityController struct {
	service *services.PersonalityService
	logger  *zap.Logger
}

func NewPersonalityController(service *services.PersonalityService, logger *zap.Logger) *PersonalityController {
	return &PersonalityController{service: service, logger: logger}
}

func (c *PersonalityController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

func (c *PersonalityController) CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	if err := json.NewDecoder(r.Body).Decode(&personality); err != nil {
		c.logger.Error("Failed to decode request body", zap.Error(err))
		HandleControllerError(w, errs.NewCustomError("InvalidRequestBody", http.StatusBadRequest, "invalid request body"))
		return
	}

	if err := personality.Validate(); err != nil {
		c.logger.Error("Invalid personality data", zap.Error(err))
		HandleControllerError(w, err)
		return
	}

	personality, err := c.service.CreatePersonality(personality)
	if err != nil {
		c.logger.Error("Failed to create personality", zap.Error(err))
		HandleControllerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(personality)
}

func (c *PersonalityController) UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	if err := json.NewDecoder(r.Body).Decode(&personality); err != nil {
		c.logger.Error("Failed to decode request body", zap.Error(err))
		HandleControllerError(w, errs.NewCustomError("InvalidRequestBody", http.StatusBadRequest, "invalid request body"))
		return
	}

	personality, err := c.service.UpdatePersonality(id, personality)
	if err != nil {
		c.logger.Error("Failed to update personality", zap.Error(err))
		HandleControllerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(personality)
}

func (c *PersonalityController) DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.service.DeletePersonality(id); err != nil {
		c.logger.Error("Failed to delete personality", zap.Error(err))
		HandleControllerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Personality deleted successfully"})
}

func (c *PersonalityController) GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	personalities, err := c.service.GetAllPersonalities()
	if err != nil {
		c.logger.Error("Failed to get all personalities", zap.Error(err))
		HandleControllerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(personalities)
}

func (c *PersonalityController) GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	personality, err := c.service.GetPersonalityByID(id)
	if err != nil {
		c.logger.Error("Failed to get personality by ID", zap.Error(err))
		HandleControllerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(personality)
}
