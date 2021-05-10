package model

import "go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `gorm:"type:longtext" json:"context"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	list  []*Article
	pager *app.Pager
}
