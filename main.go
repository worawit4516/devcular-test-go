package main

import (
	"github.com/worawit4516/devcular-test-go/src/config"
	"github.com/worawit4516/devcular-test-go/src/models"
	"github.com/worawit4516/devcular-test-go/src/routes"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDatabase()

func main() {
	defer config.DisconnectDatabase(db)
	db.AutoMigrate(&models.Todo{})
	routes.Routes()
}
