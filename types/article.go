package types

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Name   string `json:"name" gorm:"name" binding:"required"`
	Author int    `json:"author" gorm:"author" binding:"required"`
}

type UpdateArticle struct {
	Name string `json:"name" gorm:"name"`
}

type CreateArticle struct {
	Name   string `json:"name" gorm:"name"`
	Author int    `json:"author" gorm:"author" binding:"required"`
}
