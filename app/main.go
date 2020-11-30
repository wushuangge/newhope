package main

import (
	log "github.com/sirupsen/logrus"
	"newhope/app/route"
	"newhope/config"
	"newhope/util"
)

func main() {
	var err error
	//初始化配置文件
	config.InitEnvConf()
	//初始化日志
	util.InitLogger()
	//开启httpserver
	err = route.StartHttpServer()
	if err != nil {
		log.Fatal("服务器启动失败:", err.Error())
		return
	}
	select {}
}
