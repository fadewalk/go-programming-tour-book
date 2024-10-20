package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fadewalk/go-programming-tour-book/blog-service/global"
	"github.com/fadewalk/go-programming-tour-book/blog-service/internal/routers"
	"github.com/fadewalk/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
)

func init() {
    err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting error: %s", err)
	}

}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.WriteTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	s := &http.Server{
		
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
