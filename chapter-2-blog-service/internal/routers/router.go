package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-programming-tour-book-exercise/chapter-2-blog-service/docs"
	v1 "go-programming-tour-book-exercise/chapter-2-blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// set swagger: http://127.0.0.1:8080/swagger/index.html
	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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
