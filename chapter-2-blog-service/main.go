package main

import (
	"github.com/gin-gonic/gin"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/routers"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting() err: %v", err)
	}
}

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
	server.ListenAndServe()
}

func setupSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}
	// todo &global.ServerSetting
	err = set.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
