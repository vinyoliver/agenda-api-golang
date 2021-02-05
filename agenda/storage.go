package agenda

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type agendaContact struct {
	gorm.Model
	Name   string `json:"name"`
	Number string `json:"number"`
	Desc   string `json:"desc"`
}

func getAllContacts(db *gorm.DB) []agendaContact {
	var contacts []agendaContact
	db.Find(&contacts)
	fmt.Println(contacts)
	return contacts
}

func findByID(db *gorm.DB, id string) (*agendaContact, error) {
	var contact agendaContact
	result := db.First(&contact, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("Not found")
	}

	return &contact, nil
}

func findByName(db *gorm.DB, name string) []agendaContact {
	var contacts []agendaContact
	db.Where("name LIKE ?", "%"+name+"%").Find(&contacts)
	return contacts
}

func setup(db *gorm.DB) {
	//FIXME: Check for a better of doing this...
	db.AutoMigrate(&agendaContact{}) //pass blank pointer to configure the tables.
}

func connect() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Can't connect to the database...")
	}
	db.LogMode(true)
	setup(db)
	return db, err
}
