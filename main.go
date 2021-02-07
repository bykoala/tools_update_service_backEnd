package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"tool_update_service/controllers"
	"tool_update_service/dao/mysql"
	"tool_update_service/logger"
	"tool_update_service/routers"
	"tool_update_service/tools/settings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: tool_update_service config.yaml")
		return
	}
	// 1.加载配置
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("配置文件加载失败")
	}
	// 2.初始化日志
	if err := logger.Init(settings.Config.LogConfig); err != nil {
		fmt.Printf("日志初始化失败！err:%v\n", err)
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	//3.初始化MySQL
	if err := mysql.Init(settings.Config.MySQLConfig); err != nil {
		zap.L().Debug("MySQL数据库初始化失败")
	}
	defer mysql.Close() // 程序退出关闭数据库连接

	// 4.初始化gin内置校验器使用的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Debug("初始化validator翻译器异常！")
	}

	// 5.注册路由
	app := routers.SetupRouter()
	app.Run(fmt.Sprintf(":%d", settings.Config.Port))
}
