package v1

import (
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service/validate"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/convert"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
	"log"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 新增标签
// @Produce  json
// @Param token header string true "JWT Token"
// @Param CreateTagRequest body validate.CreateTagRequest true "CreateTagRequest"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := validate.CreateTagRequest{}
	svc := service.New(c.Request.Context())
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("Create app.BindAndValid err: %v", errors)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	if err := svc.CreateTag(&param); err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce  json
// @Param token header string true "JWT Token"
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := validate.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("Delete app.BindAndValid errors: %v", errors)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	if err := svc.DeleteTag(&param); err != nil {
		global.Logger.Errorf("svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce  json
// @Param token header string true "JWT Token"
// @Param UpdateTagRequest body validate.UpdateTagRequest true "UpdateTagRequest"
// @Success 200 {array} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := validate.UpdateTagRequest{}
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	// valid the request params.
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("Update app.BindAndValid err: %v", errors)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	// svc update logic.
	if err := svc.UpdateTag(&param); err != nil {
		global.Logger.Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 获取多个标签
// @Produce json
// @Param token header string true "JWT Token"
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	log.Println("tag list...")
	param := validate.ListTagRequest{}
	response := app.NewResponse(c)

	// valid the request params.
	valid, errors := app.BindAndValid(c, &param)
	// return (bool, ValidErrors), you also can design only return ValidErrors, then if ValidErrors != nil {...}
	if !valid {
		global.Logger.Errorf("app.BindAndValid errors: %v", errors)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	totalRows, err := svc.CountTag(&validate.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	pager := app.Pager{
		Page:      app.GetPage(c),
		PageSize:  app.GetPageSize(c),
		TotalRows: totalRows,
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		// 1. err -> log
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		// 2. err -> response
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}
