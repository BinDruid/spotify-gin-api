package models

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  uint   `json:"price"`
}
