package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	MySQL MySQLConfig
	Redis cache.CacheConf
	Cors  CORS `json:",optional"`
}

// CORS 跨域配置
type CORS struct {
	Mode      string          `json:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `json:"whitelist" yaml:"whitelist"`
}

// CORSWhitelist 跨域白名单
type CORSWhitelist struct {
	AllowOrigin      string `json:"allow-origin" yaml:"allow-origin"`
	AllowMethods     string `json:"allow-methods" yaml:"allow-methods"`
	AllowHeaders     string `json:"allow-headers" yaml:"allow-headers"`
	ExposeHeaders    string `json:"expose-headers" yaml:"expose-headers"`
	AllowCredentials bool   `json:"allow-credentials" yaml:"allow-credentials"`
}

// MySQLConfig MySQL 连接配置
type MySQLConfig struct {
	DSN             string // 数据库连接字符串，如 root:123456@tcp(127.0.0.1:3306)/zero-admin?charset=utf8mb4&parseTime=True&loc=Local
	MaxIdleConns    int    // 最大空闲连接数
	MaxOpenConns    int    // 最大打开连接数
	ConnMaxLifetime int    // 连接最长复用时间（秒）
}
