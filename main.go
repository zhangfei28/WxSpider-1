package main

import (
	"WxSpider/models"
	"WxSpider/pkg/logging"
	"WxSpider/pkg/setting"
	"WxSpider/router"
	"fmt"
	"net/http"
)

//初始化统一控制方法
func serverInit() {
	setting.InitSetUp("conf/app.ini")
	logging.InitSetUp()
	models.InitSetUp()
}

//销毁统一控制
func serverDestory() {
	models.CloseDb()
}

func main() {
	serverInit()
	defer serverDestory()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerConfig.HttpPort),
		Handler:        router.InitRouter(),
		ReadTimeout:    setting.ServerConfig.ReadTimeout,
		WriteTimeout:   setting.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
