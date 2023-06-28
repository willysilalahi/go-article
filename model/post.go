package model

import "time"

type Post struct {
	ID          int
	Title       string
	Description string `gorm:"type:text"`
	Status      int
	AdminID     int
	Admin       Admin
	Comments    []Comment
	Tags        []Tag `gorm:"many2many:post_tags"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
