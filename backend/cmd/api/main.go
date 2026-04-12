package main

import (
	"fmt"
	"log"

	"github.com/sapienfrom2000s/trident/backend/internal/core/models"
	"github.com/sapienfrom2000s/trident/backend/internal/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Connecting with Sqlite DB...")
	db, err := gorm.Open(sqlite.Open("trident.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected with Sqlite DB!")
	fmt.Println("Running Migrations...")
	db.AutoMigrate(&models.Event{}, &models.HeartBeat{}, &models.Job{})
	fmt.Println("Ran Migrations!")
	fmt.Println("Starting Server...")
	server.Main()
}
