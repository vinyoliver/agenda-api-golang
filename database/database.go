package database

import (
	"agenda-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setup(db *gorm.DB) {
	//FIXME: Check for a better of doing this...
	db.AutoMigrate(&models.Contact{}) //pass blank pointer to configure the tables.
}

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Can't connect to the database...")
	}
	db.LogMode(true)
	setup(db)
	return db, err
}
