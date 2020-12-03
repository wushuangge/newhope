package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"newhope/app/route"
	"newhope/app/store/mongodb"
	"newhope/config"
	"newhope/util"
)

func main() {
	var err error
	//初始化配置文件
	config.InitEnvConf()
	//初始化日志
	util.InitLogger()
	//初始化mongodb
	err = mongodb.InitMongoDB()
	if err != nil {
		log.Fatal("服务器启动失败:", err.Error())
		return
	}
	//开启httpserver
	err = route.StartHttpServer()
	if err != nil {
		log.Fatal("服务器启动失败:", err.Error())
		return
	}
	fmt.Println("http服务启动成功！！！")
	select {}
}
