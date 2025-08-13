package main

import (
	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/app"
	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/config"
	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/models"
	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/pkg/common"
)

func main() {
	config.LoadConfig([]string{"./internal/config"}, "config")
	db, err := common.ConnectDBWithConfig()
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	db.AutoMigrate(&models.State{})

	app.Start(db)
}
