package controllers

import (
	"agenda-api/database"
	"agenda-api/models"
	"agenda-api/repository"
	"agenda-api/responses"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetContacts(w http.ResponseWriter, r *http.Request) {
	db, _ := database.Connect()
	defer db.Close()
	var result []models.Contact
	if r.FormValue("name") != "" {
		result = repository.FindByName(db, r.FormValue("name"))
	} else {
		result = repository.GetAllContacts(db)
	}
	json.NewEncoder(w).Encode(result)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, _ := database.Connect()
	defer db.Close()
	result, err := repository.FindById(db, params["id"])
	if err != nil {
		responses.JSON(w, http.StatusNotFound, nil)
		return
	}
	responses.JSON(w, http.StatusOK, result)
}

func AddContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, _ := database.Connect()
	defer db.Close()
	db.Create(&contact)
	responses.JSON(w, http.StatusCreated, contact)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, _ := database.Connect()
	defer db.Close()
	result := db.Delete(&models.Contact{}, params["id"])
	if result.RowsAffected == 0 {
		responses.ERROR(w, http.StatusNotFound, errors.New("Contact not found"))
		return
	}
	var contacts []models.Contact
	db.Find(&contacts)
	json.NewEncoder(w).Encode(contacts)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, _ := database.Connect()
	defer db.Close()

	var updated models.Contact
	_ = json.NewDecoder(r.Body).Decode(&updated)

	var contact models.Contact
	db.First(&contact, params["id"])
	//TODO: set other fields....
	contact.Name = updated.Name
	db.Save(&contact)
	json.NewEncoder(w).Encode(&contact)

}
