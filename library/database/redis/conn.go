package redis

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/25 12:24
 * @Desc:
 */

type cache struct {
	pool map[string]*redis.Client
	lock sync.RWMutex
}

var Cache = &cache{
	pool: make(map[string]*redis.Client),
}

// LoadRedis 加载指定Redis
func LoadRedis(name string, config Config) {
	Cache.lock.Lock()
	defer Cache.lock.Unlock()
	Cache.pool[name] = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Pass,
		DB:       config.DB, // use default DB
	})
}

func GetRedis(name string) *redis.Client {
	Cache.lock.Lock()
	defer Cache.lock.Unlock()
	return Cache.pool[name]
}
