package models

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	ID   string `json: "id"`
	Code string `json: "code"`
	Name string `json: "name"`
}
