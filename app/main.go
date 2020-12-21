package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"newhope/app/route"
	"newhope/app/store/mongodb"
	"newhope/config"
	"newhope/util"
	"os"
)

func init()  {
	//初始化配置文件
	config.InitEnvConf()
	path := config.GetImagesPath()
	if !isDir(path) {
		fmt.Println(path, "不是一个目录, 正在创建目录", "...")
		err := os.Mkdir(path, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(path, "创建目录成功")
	}
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func main() {
	var err error
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