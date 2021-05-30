package infrastructure

import (
	"github.com/jinzhu/gorm"
)

func autoMigration(db *gorm.DB) {
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
}
