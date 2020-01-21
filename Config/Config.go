package Config


import (
"fmt"
"github.com/spf13/viper"
)

func init() {
	projectName := "casbin"
	getConfig(projectName)
}

func getConfig(projectName string) {
	// 配置文件的名称（不带扩展名）
	// name of config file (without extension)
	viper.SetConfigName("config")

	// 可选择在工作目录中查找配置
	// optionally look for config in the working directory
	viper.AddConfigPath(".")

	// 多次调用以添加许多搜索路径
	// call multiple times to add many search paths
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))

	// 查找配置文件的路径
	// path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s", projectName))

	// 查找并阅读配置文件
	// Find and read the config file
	err := viper.ReadInConfig()

	// 处理读取配置文件的错误
	// Handle errors reading the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}