package redis

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/11/25 12:22
 * @Desc:
 */

type Config struct {
	Addr        string // 地址
	Pass        string
	DB          int
	MaxIdle     int // 最大闲置连接数
	MaxActive   int // 最大连接数
	IdleTimeout int // 闲置连接有效期
}
