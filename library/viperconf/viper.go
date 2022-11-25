package viperconf

import (
	"fmt"

	"github.com/spf13/viper"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 22:05
 * @Desc:
 */

func LoadConfig(name string, v interface{}) {
	// viper.AddConfigPath(".")           // 还可以在工作目录中查找配置
	viper.SetConfigFile(name)   // 指定配置文件路径(这一句跟下面两行合起来表达的是一个意思)
	err := viper.ReadInConfig() // 配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&v)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
