package service

import (
	"net/http"
	"strconv"

	"src/github.com/mustafaakilll/rest_api/database"
	. "src/github.com/mustafaakilll/rest_api/types"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type ApiServer struct {
	storage database.DatabaseOperations
}

func NewApiServer(storage database.DatabaseOperations) *ApiServer {
	return &ApiServer{
		storage: storage,
	}
}

func (a ApiServer) HandleGetArticles(c *gin.Context) {
	if c.Query("authorId") != "" {
		return
	}
	articles, err := a.storage.GetArticles()
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		Response{
			Result: articles,
			ResponseStaus: ResponseStatus{
				Message:   nil,
				ErrorCode: 0,
			},
		},
	)
}

func (a ApiServer) HandleGetArticleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}

	articles, err := a.storage.GetArticleById(id)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		Response{
			Result: articles,
			ResponseStaus: ResponseStatus{
				Message:   nil,
				ErrorCode: 0,
			},
		},
	)
}

func (a ApiServer) HandleGetArticleByAuthor(c *gin.Context) {
	if c.Query("authorId") == "" {
		return
	}

	id, err := strconv.Atoi(c.Query("authorId"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}

	articles, err := a.storage.GetArticleByAuthor(id)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		Response{
			Result: articles,
			ResponseStaus: ResponseStatus{
				Message:   nil,
				ErrorCode: 0,
			},
		},
	)
}

func (a ApiServer) HandleCreateArticle(c *gin.Context) {
	var newarticle Article
	err := c.BindJSON(&newarticle)

	if err != nil {
		c.JSON(
			http.StatusNotAcceptable,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}

	articleId, err := a.storage.CreateArticle(newarticle)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}

	article, err := a.storage.GetArticleById(articleId)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		Response{
			Result: article,
			ResponseStaus: ResponseStatus{
				Message:   nil,
				ErrorCode: 0,
			},
		},
	)
}

func (a ApiServer) HandleUpdateArticle(c *gin.Context) {
	var newarticle UpdateArticle

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}

	if err := c.BindJSON(&newarticle); err != nil {
		c.JSON(
			http.StatusNotAcceptable,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}

	_, err = a.storage.UpdateArticle(id, newarticle)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}

	article, err := a.storage.GetArticleById(id)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		Response{
			Result: article,
			ResponseStaus: ResponseStatus{
				Message:   nil,
				ErrorCode: 0,
			},
		},
	)
}

func (a ApiServer) HandleDeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}

	if err = a.storage.DeleteArticle(id); err != nil {
		c.JSON(
			http.StatusBadRequest,
			Response{
				Result: nil,
				ResponseStaus: ResponseStatus{
					ErrorCode: 01, Message: err,
				},
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Response{
			Result: "Article deleted",
			ResponseStaus: ResponseStatus{
				Message:   nil,
				ErrorCode: 0,
			},
		},
	)
}
