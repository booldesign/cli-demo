package config

import (
	"github.com/booldesign/cli-demo/library/database/mysql"
	"github.com/booldesign/cli-demo/library/database/redis"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 22:20
 * @Desc:
 */

type Config struct {
	App   BaseConfig
	DB    mysql.Config
	Cache redis.Config
}

type BaseConfig struct {
	Name string
	Mode string
	Log  struct {
		Path  string
		Level string
	}
}
