package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID          string `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Sequence    int    `json: "sequence"`
}
