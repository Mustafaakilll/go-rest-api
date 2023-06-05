package database

import (
	"database/sql"
	"fmt"

	. "src/github.com/mustafaakilll/rest_api/types"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	db *sql.DB
}

func ConnectDB() (*SqliteDB, error) {
	db, err := sql.Open("sqlite3", "./article.db")

	if err != nil {
		return &SqliteDB{}, err
	}

	return &SqliteDB{db: db}, nil
}

type DatabaseOperations interface {
	GetArticles() ([]*Article, error)
	GetArticleById(int) (*Article, error)
	GetArticleByAuthor(int) ([]*Article, error)
	CreateArticle(Article) (int, error)
	UpdateArticle(int, UpdateArticle) (int, error)
	DeleteArticle(int) error
}

func (s SqliteDB) Init() error {
	return s.createArticleTable()
}

func (s SqliteDB) createArticleTable() error {
	query := `create table if not exists articles (
		id integer primary key,
		name text,
		author integer
	);`

	_, err := s.db.Exec(query)
	return err
}

func (s SqliteDB) GetArticles() ([]*Article, error) {
	rows, err := s.db.Query("Select * from articles")
	if err != nil {
		return nil, err
	}

	articles := []*Article{}
	for rows.Next() {
		article, err := scanIntoArticle(rows)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (s SqliteDB) GetArticleById(id int) (*Article, error) {
	rows, err := s.db.Query("Select * from articles where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		article, err := scanIntoArticle(rows)
		if err != nil {
			return nil, err
		}
		return article, nil
	}
	return nil, fmt.Errorf("article %v not found", id)
}

func (s SqliteDB) GetArticleByAuthor(authorId int) ([]*Article, error) {
	rows, err := s.db.Query("Select * from articles where author = $1", authorId)
	if err != nil {
		return nil, err
	}

	articles := []*Article{}
	for rows.Next() {
		article, err := scanIntoArticle(rows)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (s SqliteDB) CreateArticle(newarticle Article) (int, error) {
	rows, err := s.db.Exec("Insert into articles(name,author) values ($1, $2)", newarticle.Name, newarticle.Author)
	if err != nil {
		return 0, err
	}

	id, _ := rows.LastInsertId()

	return int(id), nil
}

func (s SqliteDB) UpdateArticle(id int, newarticle UpdateArticle) (int, error) {
	rows, err := s.db.Exec("update articles set name=$1 where id=$2", newarticle.Name, id)
	if err != nil {
		return 0, err
	}

	affectedRows, _ := rows.RowsAffected()

	return int(affectedRows), nil
}

func (s SqliteDB) DeleteArticle(id int) error {
	_, err := s.db.Exec("delete from articles where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func scanIntoArticle(rows *sql.Rows) (*Article, error) {
	article := new(Article)
	err := rows.Scan(
		&article.ID,
		&article.Name,
		&article.Author,
	)
	return article, err
}
