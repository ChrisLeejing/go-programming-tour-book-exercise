package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-programming-tour-book-exercise/chapter-2-blog-service/docs"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/middleware"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/routers/api"
	v1 "go-programming-tour-book-exercise/chapter-2-blog-service/internal/routers/api/v1"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/limiter"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     1,
	Quantum:      1,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.Translations())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Tracing())

	// set swagger: http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()
	auth := v1.NewAuth()

	// upload file
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.POST("/upload/files", upload.UploadFiles)

	// UploadSavePath: storage/uploads
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// auth: token
	r.POST("/auth", auth.GetAuth)

	apiV1 := r.Group("api/v1")
	// add middleware token
	apiV1.Use(middleware.JWT())
	{
		// tag: CRUD
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.GET("/tags", tag.List)
		apiV1.GET("/tags/:id", tag.Get)

		// article: CRUD
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.POST("/articles/:id", article.Update)
		apiV1.GET("articles/:id", article.Get)
		apiV1.GET("/articles", article.List)

	}

	return r
}
