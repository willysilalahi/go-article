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

	router.GET("/post", postHandler.GetAllPostHandler)

	router.Run()
}
