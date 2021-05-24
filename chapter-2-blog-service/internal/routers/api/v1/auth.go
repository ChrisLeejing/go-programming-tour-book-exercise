package v1

import (
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service/validate"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
)

type Auth struct {
}

func NewAuth() Auth {
	return Auth{}
}

// @Summary 获取认证token
// @Produce  json
// @Param GetAuthRequest body validate.GetAuthRequest true "GetAuthRequest"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [post]
func (a Auth) GetAuth(c *gin.Context) {
	// panic("模拟邮件发送错误") // 详见告警邮件截图: ./docs/email-panic/email-panic.png
	param := validate.GetAuthRequest{}
	response := app.NewResponse(c)

	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errors)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	auth, err := svc.GetAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(auth.AppKey, auth.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenError)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
	return
}
