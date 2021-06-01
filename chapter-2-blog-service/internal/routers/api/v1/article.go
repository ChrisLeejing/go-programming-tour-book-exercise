package v1

import (
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service/validate"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
	"log"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// @Summary 新增文章
// @Tags Article 文章
// @Produce  json
// @Param token header string true "JWT Token"
// @Param CreateArticleRequest body validate.CreateArticleRequest true "CreateArticleRequest"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) CreateArticle(c *gin.Context) {
	// global.Logger.Errorf("svc.dao.CreateArticle err: %v", errcode.ErrorCreateArticleFail)
	param := validate.CreateArticleRequest{}
	response := app.NewResponse(c)

	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errors)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	createArticleResponse, err := svc.CreateArticle(&param)
	if err != nil {
		return
	}

	response.ToResponse(createArticleResponse)
}

func (a Article) DeleteArticle(c *gin.Context) {
}

func (a Article) UpdateArticle(c *gin.Context) {
}

func (a Article) GetArticle(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError) // test the standard response.
	log.Println("article get...")
	return
}
func (a Article) ListArticle(c *gin.Context) {
}
