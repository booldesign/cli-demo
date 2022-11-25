package svc

import (
	"github.com/booldesign/cli-demo/daemon/internal/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/24 22:20
 * @Desc:
 */

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config, db *gorm.DB, r *redis.Client) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     db,
		Redis:  r,
	}
}
