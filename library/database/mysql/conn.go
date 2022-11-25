package mysql

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/25 13:38
 * @Desc: 数据库加载
 */

type database struct {
	pool map[string]*gorm.DB
	lock sync.RWMutex
}

var AppDatabase = &database{
	pool: make(map[string]*gorm.DB),
}

// Assign 加载数据库
func Assign(name string, config Config, plugins ...gorm.Plugin) error {
	myLogger := logger.Default
	if config.Mode {
		myLogger = myLogger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(config.Url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 全局禁用表名复数
			TablePrefix:   config.Prefix,
		},
		Logger: myLogger, // 启用内置Logger，显示详细日志
	})
	if err != nil {
		return fmt.Errorf("init db[%s] connection error: %s", name, err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("init db[%s] connection error: %s", name, err.Error())
	}
	if err = sqlDB.Ping(); err != nil {
		return fmt.Errorf("init db[%s] ping error: %s", name, err.Error())
	}
	sqlDB.SetMaxIdleConns(config.MaxIdle)
	sqlDB.SetMaxOpenConns(config.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)
	for _, plugin := range plugins {
		err := db.Use(plugin)
		if err != nil {
			return err
		}
	}

	AppDatabase.lock.Lock()
	AppDatabase.pool[name] = db
	AppDatabase.lock.Unlock()

	log.Printf("db[%s] init success\n", name)

	return nil
}

func GetGormDB(ctx context.Context, name string) (*gorm.DB, error) {
	AppDatabase.lock.RLock()
	defer AppDatabase.lock.RUnlock()
	db, ok := AppDatabase.pool[name]
	if !ok {
		return nil, fmt.Errorf("db[%s] is not initial", name)
	}
	return db.WithContext(ctx), nil
}

func (db *database) CloseDatabase() []error {
	errs := make([]error, 0)
	for name, db := range AppDatabase.pool {
		sqlDB, err := db.DB()
		if err != nil {
			errs = append(errs, fmt.Errorf("db[%s] close err: %s\n", name, err.Error()))
			continue
		}
		if err := sqlDB.Close(); err != nil {
			errs = append(errs, fmt.Errorf("db[%s] close err: %s\n", name, err.Error()))
			continue
		}
		log.Printf("db[%s] close success\n", name)
	}
	return errs
}

// CloseHook 优雅关闭
func CloseHook() func() {
	return func() {
		if errs := AppDatabase.CloseDatabase(); len(errs) > 0 {
			for _, v := range errs {
				log.Println(v.Error())
			}
		}
	}
}
