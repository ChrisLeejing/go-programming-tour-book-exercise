package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ""
		ecode := errcode.Success

		if t, ok := c.GetQuery("token"); ok {
			token = t
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
