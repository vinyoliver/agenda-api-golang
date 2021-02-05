package agenda

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Service represents our service structure
type Service struct {
}

// NewService returns a newly configured instance of service
func NewService(db *gorm.DB) *Service {
	return &Service{}
}

func (svc *Service) getContacts(w http.ResponseWriter, r *http.Request) {
	db, _ := connect()
	defer db.Close()
	var result []agendaContact
	if r.FormValue("name") != "" {
		result = findByName(db, r.FormValue("name"))
	} else {
		result = getAllContacts(db)
	}
	json.NewEncoder(w).Encode(result)
}

func (svc *Service) getByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, _ := connect()
	defer db.Close()
	result, err := findByID(db, params["id"])
	if err != nil {
		JSON(w, http.StatusNotFound, nil)
		return
	}
	JSON(w, http.StatusOK, result)
}

func (svc *Service) addContact(w http.ResponseWriter, r *http.Request) {
	var contact agendaContact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, _ := connect()
	defer db.Close()
	db.Create(&contact)
	JSON(w, http.StatusCreated, contact)
}

func (svc *Service) deleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, _ := connect()
	defer db.Close()
	result := db.Delete(&agendaContact{}, params["id"])
	if result.RowsAffected == 0 {
		ERROR(w, http.StatusNotFound, errors.New("Contact not found"))
		return
	}
	var contacts []agendaContact
	db.Find(&contacts)
	json.NewEncoder(w).Encode(contacts)
}

func (svc *Service) updateContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, _ := connect()
	defer db.Close()

	var updated agendaContact
	_ = json.NewDecoder(r.Body).Decode(&updated)

	var contact agendaContact
	db.First(&contact, params["id"])
	//TODO: set other fields....
	contact.Name = updated.Name
	db.Save(&contact)
	json.NewEncoder(w).Encode(&contact)

}

// JSON writes out a standard JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	}

}

// ERROR displays a json response message with an error
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
