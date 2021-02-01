package basic

import (
	"grocery/user-service/basic/config"
	"grocery/user-service/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
