package api

import (
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/convert"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/upload"
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
