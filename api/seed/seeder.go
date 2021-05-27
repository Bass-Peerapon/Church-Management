package seed

import (
	"log"

	"github.com/Church-Management/api/models"
	"gorm.io/gorm"
)


func Load(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}