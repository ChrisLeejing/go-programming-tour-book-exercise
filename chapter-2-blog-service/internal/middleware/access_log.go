package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

// override the Write() of gin.ResponseWriter, http.ResponseWriter
func (w *AccessLogWriter) Write(data []byte) (n int, err error) {
	if nn, err := w.buf.Write(data); err != nil {
		return nn, err
	}

	return w.ResponseWriter.Write(data)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := &AccessLogWriter{buf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		begin := time.Now().Unix()
		c.Next()
		end := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": writer.buf.String(),
		}

		s := "access log: request method: %s, status: %d, begin time: %d, end time: %d"
		global.Logger.WithFields(fields).Infof(c, s, c.Request.Method, c.Status, begin, end)
	}
}
