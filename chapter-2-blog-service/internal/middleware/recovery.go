package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/email"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				// QQ 邮箱-设置-账户-POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV 服务选项中
				// 将POP3/SMTP 服务和IMAP/SMTP 服务开启，然后根据所获取的 SMTP 账户密码进行设置即可，另外 SSL 是默认开启的。
				// 告警邮件截图: ./docs/email-panic/email-panic.png
				err := defaultMailer.SendEmail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf("mail.SendMail err: %v", err)
				}
				c.Abort()
			}
		}()

		c.Next()
	}
}
