package model

import "go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	list  []*Tag
	pager *app.Pager
}
