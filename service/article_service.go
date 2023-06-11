package service

import (
	"net/http"
	"strconv"

	"src/github.com/mustafaakilll/rest_api/database"
	. "src/github.com/mustafaakilll/rest_api/types"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type ArticleService struct {
	storage database.DatabaseOperations
}

func NewArticleService(storage database.DatabaseOperations) *ArticleService {
	return &ArticleService{
		storage: storage,
	}
}

func (a ArticleService) HandleGetArticles(c *gin.Context) {
	if c.Query("authorId") != "" {
		return
	}
	articles, err := a.storage.GetArticles()
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	c.JSON(SuccessResponse(articles))
}

func (a ArticleService) HandleGetArticleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	article, err := a.storage.GetArticleById(id)
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	c.JSON(SuccessResponse(article))
}

func (a ArticleService) HandleGetArticleByAuthor(c *gin.Context) {
	if c.Query("authorId") == "" {
		return
	}

	id, err := strconv.Atoi(c.Query("authorId"))
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	articles, err := a.storage.GetArticleByAuthor(id)
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	c.JSON(SuccessResponse(articles))
}

func (a ArticleService) HandleCreateArticle(c *gin.Context) {
	var newarticle CreateArticle
	if err := c.BindJSON(&newarticle); err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	articleId, err := a.storage.CreateArticle(newarticle)
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusNotAcceptable))
		return
	}

	article, err := a.storage.GetArticleById(articleId)
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	c.JSON(SuccessResponse(article))
}

func (a ArticleService) HandleUpdateArticle(c *gin.Context) {
	var newarticle UpdateArticle

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	if err := c.BindJSON(&newarticle); err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusNotAcceptable))
		return
	}

	err = a.storage.UpdateArticle(id, newarticle)
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}

	article, err := a.storage.GetArticleById(id)
	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusNotFound))
		return
	}
	c.JSON(SuccessResponse(article))
}

func (a ArticleService) HandleDeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	if err = a.storage.DeleteArticle(id); err != nil {
		c.JSON(ErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(SuccessResponse("Article deleted successfully"))
}
