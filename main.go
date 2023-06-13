package main

import (
	"log"
	"net/http"
	"src/github.com/mustafaakilll/rest_api/database"
	"src/github.com/mustafaakilll/rest_api/middleware"
	"src/github.com/mustafaakilll/rest_api/service"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	articleService := service.NewArticleService(db)
	userService := service.NewUserService(db)

	r := gin.New()
	r.Use(logger.SetLogger())

	r.GET("/ping", ping)
	r.GET("/articles", articleService.HandleGetArticles, articleService.HandleGetArticleByAuthor)
	r.GET("/articles/:id", articleService.HandleGetArticleById)
	r.PUT("/articles/:id", middleware.Auth(), articleService.HandleUpdateArticle)
	r.POST("/articles", middleware.Auth(), articleService.HandleCreateArticle)
	r.DELETE("/articles/:id", articleService.HandleDeleteArticle)

	r.POST("/register", userService.HandleRegisterUser)
	r.POST("/login", userService.HandleLoginUser)

	http.ListenAndServe(":3000", r)
}

func ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
