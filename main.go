package main

import (
	"go-article/database"
	"go-article/handler"
	"go-article/repository"
	"go-article/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	/*
		Your Task
		- Get All Post*
		- Get Post By Id*
		- Create Post
		- Update Post
		- Delete Post
		- Get Comments By Post ID
		- Create Comment
		- Edit Comment
		- Delete Comment
	*/

	dsn := "root:@tcp(127.0.0.1:3306)/go_article?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database error connection")
	}
	database.Migrate(db)
	database.Seeding(db)

	postRepository := repository.NewRepository(db)
	postService := service.NewService(postRepository)
	postHandler := handler.NewHandler(postService)

	router := gin.Default()

	router.GET("/article", postHandler.GetAllPostHandler)
	router.GET("/article/:id", postHandler.GetSinglePostHandler)

	router.Run()
}
