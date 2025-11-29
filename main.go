package main

import (
	"jachow/code1024/config"
	"jachow/code1024/dao/mysql"
	"jachow/code1024/dao/redis"
	"jachow/code1024/logger"
	"jachow/code1024/pkg"
	"jachow/code1024/router"
)

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
