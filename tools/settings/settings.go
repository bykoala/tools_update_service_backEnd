package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"tool_update_service/models"
)

var Config = new(models.AppConfig)

// 初始化日志
func Init(filePath string) (err error) {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件失败\n;err:%v\n", err)
	}
	if err = viper.Unmarshal(Config); err != nil {
		fmt.Printf("viper unmarshal失败")
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err = viper.Unmarshal(Config); err != nil {
			fmt.Printf("viper unmarshal失败")
		}
		fmt.Printf("配置文件发生了改变...\n")
	})
	return
}
