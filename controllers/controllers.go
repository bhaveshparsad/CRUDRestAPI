package controllers

import (
	"encoding/json"
	"errors"
	"CRUDRestAPI/database"
	"CRUDRestAPI/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func respondError(w http.ResponseWriter, code int, message string) {
	respondJson(w, code, map[string]string{"error": message})
}

func respondJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var contacts []model.ContactService
	result := database.Connector.Find(&contacts)
	if result.Error != nil {
		respondError(w, http.StatusInternalServerError, "not found")
	}
	respondJson(w, http.StatusOK, contacts)
}

func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid person ID")

	}
	key := vars["id"]

	var contact model.ContactService
	result := database.Connector.First(&contact, key)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			respondError(w, http.StatusNotFound, "error occurred while fetching person with given id")
		} else {
			respondError(w, http.StatusInternalServerError, err.Error())
		}

	}
	respondJson(w, http.StatusOK, contact)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var contact model.ContactService
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&contact); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request")
	}

	result := database.Connector.Create(contact)
	if result.Error != nil && result.RowsAffected != 1 {
		log.Fatal(result.Error)
		respondError(w, http.StatusNotFound, "error occurred while creating a new person")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	var contact model.ContactService
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&contact); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request")
	}

	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid contact ID")
		return
	}

	var value model.ContactService
	result := database.Connector.First(&value, "id=?", contact.ID)
	if result.Error != nil {
		log.Fatal(result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			respondError(w, http.StatusNotFound, "error ocurred while updating contact with given id")
		} else {
			respondError(w, http.StatusInternalServerError, err.Error())
		}
	}

	value.FirstName = contact.FirstName
	value.LastName = contact.LastName
	re := database.Connector.Save(&value)
	if re.RowsAffected != 1 {
		respondError(w, http.StatusInternalServerError, "error in updating contact")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(value)
}

func DeletePersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid contact ID")
		return
	}
	key := vars["id"]

	var contact model.ContactService

	id, _ := strconv.ParseInt(key, 10, 64)
	result := database.Connector.Where("id = ?", id).Delete(&contact)
	if result.RowsAffected != 1 {
		respondError(w, http.StatusInternalServerError, "error ocurred while deleting contact")
	}

	w.WriteHeader(http.StatusNoContent)
}
