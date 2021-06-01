package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
)

type ArticleTag struct {
	*Model
	ArticleID uint32 `json:"article_id"`
	TagID     uint32 `json:"tag_id"`
}

func (ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (at ArticleTag) CreateBatchArticleTags(db *gorm.DB, tagIds []uint32) (articleTags []ArticleTag, err error) {
	articleTags = batchArticleTags(at.ArticleID, tagIds)
	sqlStr := "insert into blog_article_tag(article_id, tag_id) values "
	valueStr := ""
	for i, tagId := range tagIds {
		valueStr += fmt.Sprint("(", at.ArticleID, ",", tagId, ")")
		if i < len(tagIds)-1 {
			valueStr += ","
		}
	}
	sqlStr += valueStr
	err = db.Exec(sqlStr).Error
	if err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (at ArticleTag) CreateArticleTags(db *gorm.DB, tagIds []uint32) (articleTags []ArticleTag, err error) {
	articleTags = batchArticleTags(at.ArticleID, tagIds)
	for _, articleTag := range articleTags {
		err = db.Create(&articleTag).Error
		if err != nil {
			global.Logger.Error(errcode.ErrorCreateArticleTagsFail)
			return nil, err
		}
	}

	return articleTags, nil
}

func batchArticleTags(articleId uint32, tagIds []uint32) (articleTags []ArticleTag) {
	for _, tagId := range tagIds {
		articleTag := ArticleTag{
			ArticleID: articleId,
			TagID:     tagId,
		}

		articleTags = append(articleTags, articleTag)
	}

	return articleTags
}
