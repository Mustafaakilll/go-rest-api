package database

import (
	. "src/github.com/mustafaakilll/rest_api/types"

	_ "github.com/mattn/go-sqlite3"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDB struct {
	db *gorm.DB
}

func ConnectDB() (*SqliteDB, error) {
	db, err := gorm.Open(sqlite.Open("article.db"), &gorm.Config{})

	if err != nil {
		return &SqliteDB{}, err
	}

	return &SqliteDB{db: db}, nil
}

type DatabaseOperations interface {
	GetArticles() ([]*Article, error)
	GetArticleById(int) (*Article, error)
	GetArticleByAuthor(int) ([]*Article, error)
	CreateArticle(CreateArticle) (int, error)
	UpdateArticle(int, UpdateArticle) (*Article, error)
	DeleteArticle(int) error
}

func (s SqliteDB) Init() error {
	return s.createArticleTable()
}

func (s SqliteDB) createArticleTable() error {
	err := s.db.AutoMigrate(&Article{})
	return err

}

func (s SqliteDB) GetArticles() ([]*Article, error) {
	var articles []*Article
	result := s.db.Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func (s SqliteDB) GetArticleById(id int) (*Article, error) {
	var article *Article
	result := s.db.First(&article, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return article, nil
}

func (s SqliteDB) GetArticleByAuthor(authorId int) ([]*Article, error) {
	var articles []*Article
	result := s.db.Find(&articles, "author=?", authorId)
	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}

func (s SqliteDB) CreateArticle(newarticle CreateArticle) (int, error) {
	var article Article = Article{
		Name:   newarticle.Name,
		Author: newarticle.Author,
	}
	result := s.db.Create(&article)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(article.ID), nil
}

func (s SqliteDB) UpdateArticle(id int, newarticle UpdateArticle) (*Article, error) {
	var article Article
	if err := s.db.Where("id = ?", id).First(&article).Error; err != nil {
		return &Article{}, err
	}

	result := s.db.Model(article).Updates(newarticle)
	if result.Error != nil {
		return &Article{}, result.Error
	}

	return &article, nil
}

func (s SqliteDB) DeleteArticle(id int) error {
	result := s.db.Delete(&Article{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
