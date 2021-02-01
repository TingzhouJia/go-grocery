package main

import (
	"github.com/jinzhu/gorm"
	"grocery/handler"
	"grocery/repository"
	"grocery/service"
)

var (
	DB              *gorm.DB
	UserBaseHandler     handler.UserHandler
	ProductBaseHandler handler.ProductHandler
)

func initHandler()  {
	UserBaseHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: DB,
			},
		},
	}
	ProductBaseHandler=handler.ProductHandler{
		ProductSrv: &service.ProductService{
			Repo: &repository.ProductRepository{
				DB: DB,
			},
		},
	}
}

func init()  {
	initHandler()
}

