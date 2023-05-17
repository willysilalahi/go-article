package database

import (
	"fmt"
	"go-article/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		model.Tag{},
		model.Admin{},
		model.Post{},
		model.Comment{},
	)
	fmt.Println("Migrate successfully")
}
