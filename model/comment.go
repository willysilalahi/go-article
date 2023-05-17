package model

import "time"

type Comment struct {
	ID        int    `json:"id"`
	Comment   string `json:"comment"`
	PostID    int    `json:"post_id"`
	Post      Post
	AdminID   int `json:"admin_id"`
	Admin     Admin
	CreatedAt time.Time
	UpdatedAt time.Time
}
