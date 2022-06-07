package migrations

import (
	"github.com/Dihh-barret/APIBooks/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
