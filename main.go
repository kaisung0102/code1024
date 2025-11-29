package main

import (
	"jachow/code1024/config"
	"jachow/code1024/dao/mysql"
	"jachow/code1024/dao/redis"
	"jachow/code1024/logger"
	"jachow/code1024/pkg"
	"jachow/code1024/router"


)

// @title Code1024
// @version 1.0
// @description Code1024是一个基于Golang的论坛项目
// @termsOfService http://swagger.io/terms/

// @contact.name jachow
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1

func main() {
	
	// 读取配置文件
	config.InitConfig("./config/config.yaml")
	// 初始化日志
	logger.InitLogger(&config.Conf.Log)
	// 初始化数据库
	mysql.InitMysql()
	// 初始化Redis
	redis.InitRedis(&config.Conf.Redis)
	// 初始化路由
	r := router.Routers()
	// 初始化雪花算法
	pkg.InitID(config.Conf.StartTime, config.Conf.MachineID)
	// 启动路由
	err := r.Run(config.Conf.Port)
	if err != nil {
		panic(err)
	}

}
