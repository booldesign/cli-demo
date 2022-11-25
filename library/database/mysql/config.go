package mysql

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/9/26 15:37
 * @Desc:
 */

// Config 数据库配置
type Config struct {
	Url         string `json:"url"`         // 地址
	Mode        bool   `json:"mode"`        // 运行模式 TRUE开启调试模式
	MaxIdle     int    `json:"maxIdle"`     // 最大闲置连接数
	MaxOpen     int    `json:"maxOpen"`     // 最大连接数
	MaxLifetime int    `json:"maxLifetime"` //最大周期
	Prefix      string `json:"prefix"`      //表前缀
}
