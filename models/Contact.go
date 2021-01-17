package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Number string `json:"number"`
	Desc   string `json:"desc"`
}
