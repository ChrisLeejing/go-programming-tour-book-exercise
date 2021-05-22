package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/convert"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/upload"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/util"
	"net/http"
	"strings"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

// @Summary 文件上传
// @Description
// @Tags file
// @Accept multipart/form-data
// @version 1.0
// @Param type formData int true "文件类型"
// @Param file formData file true "file"
// @Produce  json
// @success 200 {string} string "成功"
// @Router /upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	// // PostForm returns the specified key from a POST urlencoded form or multipart form
	// // when it exists, otherwise it returns an empty string `("")`.
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	ctx := c.Request.Context()
	svc := service.New(ctx)
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	// svc := service.New(c.Request.Context())
	// fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{"file_access_url": fileInfo.AccessUrl})
	return
}

// curl -X POST http://localhost:8080/upload/files   -F "upload[]=@./docs/uploadfiles/1.gif"   -F "upload[]=@./docs/uploadfiles/2.jpg"   -H "Content-Type: multipart/form-data"
// http://localhost:8080/static/c4ca4238a0b923820dcc509a6f75849b.gif
// http://localhost:8080/static/c81e728d9d4c2f636f067f89cc14862c.jpg
func (u Upload) UploadFiles(c *gin.Context) {
	response := app.NewResponse(c)
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		fmt.Println(file.Filename)
		dot := strings.LastIndex(file.Filename, ".")
		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, "storage/uploads/"+util.EncodeMD5(file.Filename[:dot])+file.Filename[dot:])
		if err != nil {
			global.Logger.Errorf("svc.UploadFile err: %v", err)
			response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
			continue
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
