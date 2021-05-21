package validate

type GetAuthRequest struct {
	// https://github.com/gin-gonic/gin#model-binding-and-validation
	AppKey    string `form:"app_key" json:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" json:"app_secret" binding:"required"`
}
