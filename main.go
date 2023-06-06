package main

import (
	"log"
	"net/http"
	"src/github.com/mustafaakilll/rest_api/database"
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

	apiServer := service.NewApiServer(db)

	r := gin.New()
	r.Use(logger.SetLogger())

	r.GET("/ping", ping)
	r.GET("/articles", apiServer.HandleGetArticles, apiServer.HandleGetArticleByAuthor)
	r.GET("/articles/:id", apiServer.HandleGetArticleById)
	r.PUT("/articles/:id", apiServer.HandleUpdateArticle)
	r.POST("/articles", apiServer.HandleCreateArticle)
	r.DELETE("/articles/:id", apiServer.HandleDeleteArticle)

	http.ListenAndServe(":3000", r)
}

func ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
