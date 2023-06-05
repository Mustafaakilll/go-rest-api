package types

type Article struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author int    `json:"author"`
}

type UpdateArticle struct {
	Name string `json:"name"`
}

type CreateArticle struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}
