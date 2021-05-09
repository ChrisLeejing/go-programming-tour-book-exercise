package v1

import (
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
	"log"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Create(c *gin.Context) {

}

func (a Article) Delete(c *gin.Context) {
}

func (a Article) Update(c *gin.Context) {
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError) // test the standard response.
	log.Println("article get...")
	return
}
func (a Article) List(c *gin.Context) {
}
