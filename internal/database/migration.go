package database

import (
	user "github.com/Galbeyte1/go-user-man-ser/internal/user"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	// AutoMigrate takes comment struct defintion and defines all the columns
	if result := db.AutoMigrate(&user.User{}); result.Error != nil {
		return result.Error
	}
	return nil
}