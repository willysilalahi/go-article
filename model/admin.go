package model

type Admin struct {
	ID    int
	Name  string
	Email string
	Posts []Post
}
