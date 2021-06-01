package service

import (
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/model"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service/response"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service/validate"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
)

func (svc *Service) CreateArticle(param *validate.CreateArticleRequest) (*response.CreateArticleResponse, error) {
	var err error
	var tagNames []string
	article := &model.Article{
		Title:       param.Title,
		Desc:        param.Desc,
		CoverImgUrl: param.CoverImgUrl,
		Content:     param.Content,
		State:       param.State,
		Model:       &model.Model{CreatedBy: param.CreatedBy},
	}

	// table: blog_article, blog_article_tag.
	article, err = svc.dao.CreateArticle(article, param.TagIds)
	if err != nil {
		return nil, errcode.ErrorCreateArticleFail
	}

	// table: blog_tag
	tags, err := svc.dao.GetTagByIds(param.TagIds)
	if err != nil {
		return nil, err
	}
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	createArticleResponse := response.CreateArticleResponse{
		Article:  *article,
		TagNames: tagNames,
	}

	return &createArticleResponse, nil
}
