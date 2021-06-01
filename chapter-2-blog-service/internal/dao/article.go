package dao

import (
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/model"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
)

func (d *Dao) CreateArticle(param *model.Article, tagIds []uint32) (*model.Article, error) {
	// transaction begin.
	tx := d.engine.Begin()

	a := model.Article{
		Title:       param.Title,
		Desc:        param.Desc,
		CoverImgUrl: param.CoverImgUrl,
		Content:     param.Content,
		State:       param.State,
		Model:       &model.Model{CreatedBy: param.CreatedBy},
	}
	article, err := a.Create(tx)
	if err != nil {
		// transaction rollback.
		tx.Rollback()
		return nil, err
	}

	// table: blog_article_tag
	articleId := article.ID
	articleTag := model.ArticleTag{
		ArticleID: articleId,
	}

	// choice either of (CreateArticleTags, CreateArticleTags) is ok.
	// _, err = articleTag.CreateArticleTags(tx, tagIds)
	_, err = articleTag.CreateBatchArticleTags(tx, tagIds)
	if err != nil {
		// transaction rollback.
		tx.Rollback()
		return nil, err
	}

	// transaction commit.
	err = tx.Commit().Error
	if err != nil {
		// transaction rollback.
		tx.Rollback()
		global.Logger.Error(errcode.ErrorTransactionCommitFail)
		return nil, err
	}

	return article, nil
}
