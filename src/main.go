package main

import (
	"./common"
	"./controller"
	"./sys"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yejiansnake/go-yedb"
)

func main() {
	//获取配置
	if err := loadConfig() ; err != nil {
		fmt.Printf("loal config failed, info:%v", err)
		return
	}

	//初始化全局对象
	if err := initAppSetting() ; err != nil {
		fmt.Printf("init app setting failed, info:%v", err)
		return
	}

	//启动应用
	if err := runApp() ; err != nil {
		fmt.Printf("run app failed, info:%v", err)
		return
	}

	fmt.Printf("run app success \r\n")

	sys.AppInstance.Logger.Fatal(nil)
}

func loadConfig() error {
	fmt.Printf("config load start\r\n")
	err := sys.ConfigMgrInstance.Load()

	if err != nil {
		return err
	}

	fmt.Printf("config load finish\r\n")

	return nil
}

func initAppSetting() error {
	//设置app的全局变量
	sys.ConfigMgrInstance.Build(sys.AppInstance)

	//初始化全局数据
	if err := initDB(&sys.ConfigMgrInstance.Data); err != nil {
		return err
	}

	//初始化路由配置
	controller.CreateRoutes()

	return nil
}

func initDB(config *sys.ConfigInfo) error {
	yedb.DbConfigMgrInstance.Set(common.DB_NAME_TEST,
		&yedb.DbConfig{config.DB.Driver,
			config.DB.Addr,
			config.DB.Name,
			config.DB.User,
			config.DB.Pwd})

	return nil
}

func runApp() error {
	return sys.AppInstance.StartServer(sys.ConfigMgrInstance.GetServerConfig())
}