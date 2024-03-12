package main

import (
	"WebVideoServer/setting"
	"WebVideoServer/snowflake"
	"WebVideoServer/web/model/mysql"
	"WebVideoServer/web/route"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func initApp() error {

	fmt.Println(os.Args[1])
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		return err
	}
	// 加载 MySQL db
	if err := mysql.InitMysql(setting.Conf.MySQLConfig); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		return err
	}
	// 加载 雪花 id
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		zap.L().Fatal("init redis failed, err:", zap.Error(err))
		return err
	}
	return nil
}

func main() {

	// 检验参数是否有指定配置文件路径
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		os.Exit(1)
	}

	// 初始化各种应用
	if err := initApp(); err != nil {
		os.Exit(1)
	}

	route.OpenRoute()
}
