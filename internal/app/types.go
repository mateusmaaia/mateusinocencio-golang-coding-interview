package app

import (
	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/models"
)

type BaseResponse struct {
	Status string      `json:"status"`
	Info   interface{} `json:"info"`
}

type StatesResponse struct {
	BaseResponse
	States []models.State `json:"data"`
}
