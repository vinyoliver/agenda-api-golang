package repository

import (
	"agenda-api/models"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetAllContacts(db *gorm.DB) []models.Contact {
	var contacts []models.Contact
	db.Find(&contacts)
	fmt.Println(contacts)
	return contacts
}

func FindById(db *gorm.DB, id string) (*models.Contact, error) {
	var contact models.Contact
	result := db.First(&contact, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("Not found")
	} else {
		return &contact, nil
	}
}

func FindByName(db *gorm.DB, name string) []models.Contact {
	var contacts []models.Contact
	db.Where("name LIKE ?", "%"+name+"%").Find(&contacts)
	return contacts
}
