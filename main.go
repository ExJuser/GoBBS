package main

import (
	"GoBBS/dao/mysql"
	"GoBBS/dao/redis"
	setting "GoBBS/settings"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file. eg: gobbs config.yaml")
		return
	}
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
}
