package model

import "time"

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	Status      int    `json:"status"`
	AdminID     int    `json:"admin_id"`
	Admin       Admin
	Comments    []Comment
	Tags        []Tag     `gorm:"many2many:post_tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
