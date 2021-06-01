package response

import "go-programming-tour-book-exercise/chapter-2-blog-service/internal/model"

type CreateArticleResponse struct {
	model.Article
	TagNames []string
}
