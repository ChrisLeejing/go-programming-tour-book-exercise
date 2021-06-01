package model

import (
	"github.com/jinzhu/gorm"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
)

type Article struct {
	*Model
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	CoverImgUrl string `json:"cover_img_url"`
	Content     string `gorm:"type:longtext" json:"context"`
	State       uint8  `json:"state"`
}

func (Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	list  []*Article
	pager *app.Pager
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		global.Logger.Error(errcode.ErrorCreateArticleOnlyFail)
		return nil, err
	}

	return &a, nil
}
