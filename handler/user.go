package controller

import (
	"grocery/dto"
	"grocery/model"
	"grocery/service"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (receiver *UserHandler) GetEntity( result model.User) dto.User  {

}