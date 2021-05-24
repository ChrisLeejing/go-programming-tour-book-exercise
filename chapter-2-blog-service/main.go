package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/model"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/routers"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/logger"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/setting"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/validator"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	// set config.
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting() err: %v", err)
	}

	// set db engine.
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setupDBEngine() err: %v", err)
	}

	// set log component.
	err = setupLogger()
	if err != nil {
		log.Fatalf("init setupLogger() err: %v", err)
	}

	// set custom validator.
	err = setupValidator()
	if err != nil {
		log.Fatalf("init setupValidator() err: %v", err)
	}
}

// @title 博客系统
// @version v1.0
// @description Go 语言编程之旅: 一起用 Go 做项目
// @termsOfService https://github.com/ChrisLeejing/go-programming-tour-book-exercise
func main() {
	// set server run mode.
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	server := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	// test the log component.
	global.Logger.Infof("%s: go-programming-tour-book-exercise/%s", "coirlee", "blog-service")
	server.ListenAndServe()

}

func setupSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = set.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	// note: '=' can not be ":=", or global.DBEngine will be local variable.
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupValidator() error {
	global.Validator = validator.NewCustomValidator()
	global.Validator.Engine()

	binding.Validator = global.Validator

	return nil
}
