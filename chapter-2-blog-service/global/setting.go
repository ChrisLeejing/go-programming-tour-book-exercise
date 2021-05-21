package global

import (
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/logger"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/setting"
)

// the global variables.
var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	JWTSetting      *setting.JWTSetting
	Logger          *logger.Logger
)
