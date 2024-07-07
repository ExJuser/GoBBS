package main

import (
	"GoBBS/dao/mysql"
	"GoBBS/dao/redis"
	"GoBBS/logger"
	"GoBBS/routers"
	setting "GoBBS/settings"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func main() {
	//从命令行获取配置文件
	if len(os.Args) < 2 {
		fmt.Println("need config file. eg: gobbs config.yaml")
		return
	}

	//读取配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	//初始化日志
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer func(l *zap.Logger) {
		if err := l.Sync(); err != nil {
			fmt.Printf("zap sync failed, err:%v\n", err)
		}
	}(zap.L())

	//初始化MySQL
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	//初始化Redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	//注册路由
	r := routers.Setup()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
