package routers

import (
	"github.com/gin-gonic/gin"
	v1 "go-programming-tour-book-exercise/chapter-2-blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiV1 := r.Group("api/v1")
	{
		// tag: CRUD
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.GET("/tags/:id", tag.List)

		// article: CRUD
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.POST("/articles/:id", article.Update)
		apiV1.GET("articles/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}

	return r
}
