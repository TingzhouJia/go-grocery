package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"grocery/user-service/model/user"

	userservice "grocery/user-service/proto"
)

type UserService struct{}



var (
	userService user.Service
)
func Init() {
	var err error
	userService, err = user.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

func (e *UserService) QueryUserByName(ctx context.Context, req *userservice.Request, rsp *userservice.Response) error {

	user,err:=userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &userservice.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return nil
	}
	rsp.User=user
	rsp.Success=true

	return nil
}

