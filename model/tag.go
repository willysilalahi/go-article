package model

type Tag struct {
	ID   int    `gorm:"id"`
	Name string `gorm:"name"`
}
