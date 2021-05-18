package controller

type CreateArticleRequest struct {
	Name string `form:"name" json:"name" binding:"required,max=100"`
}

type DeleteArticleRequest struct {
	// todo 20210517
}

type UpdateArticleRequest struct {
}

type ListArticleRequest struct {
}

type CountArticleRequest struct {
}
