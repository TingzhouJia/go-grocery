package basic

import (
	"grocery/basic/config"
	"grocery/basic/db"
	"grocery/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
