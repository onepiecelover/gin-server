package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./config/") // 比如添加当前目录
	viper.SetConfigName("config")    //  设置配置文件名 (不带后缀)
	//viper.AddConfigPath("/etc/appname/")  // 第一个搜索路径
	//viper.AddConfigPath("$HOME/.appname") // 可以多次调用添加路径
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//监视配置文件，重新读取配置数据
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
