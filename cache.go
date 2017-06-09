package utils

import (
	bc "github.com/bysir-zl/bygo/cache"
	"github.com/astaxie/beego"
)

var Redis = bc.NewRedis("")

func init() {
	redisHost := beego.AppConfig.String("redis_host")
	if redisHost == "" {
		redisHost = "10.8.230.17:6379"
	}
	Redis = bc.NewRedis(redisHost)
}
