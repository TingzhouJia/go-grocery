package handler

import (
	"github.com/gin-gonic/gin"
	"grocery/dto"
	"grocery/enum"
	"grocery/model"
	"grocery/query"
	"grocery/service"
	"grocery/utils"
	"net/http"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (receiver *UserHandler) GetUser( result model.User) dto.User  {
	return dto.User{
		Id:        result.UserId,
		Key:       result.UserId,
		UserId:    result.UserId,
		NickName:  result.NickName,
		Mobile:    result.Mobile,
		Address:   result.Address,
		IsDeleted: result.IsDeleted,
		IsLocked:  result.IsLocked,
	}
}

func (h *UserHandler) GetUserInfo(c *gin.Context)  {
	baseBody:=utils.BaseReturnBody()
	userId:=c.Param("id")
	if userId!="" {
		c.JSON(http.StatusInternalServerError,gin.H{"data":baseBody})
	}
	u:=model.User{
		UserId: userId,
	}
	result,err:=h.UserSrv.Get(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": baseBody})
		return
	}
	r := h.GetUser(*result)
	entity:=utils.ReturnBody(0,0,enum.Success,r)

	c.JSON(http.StatusOK,gin.H{"data":entity})

}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	b, err := h.UserSrv.Delete(id)
	entity := utils.BaseReturnBody()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = int(enum.DeleteSuccess)
		entity.Msg = enum.DeleteSuccess.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}

func (h *UserHandler) EditUserHandler(c *gin.Context) {
	u := model.User{}
	entity := utils.BaseReturnBody()
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	b, err := h.UserSrv.Edit(u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = int(enum.Accepted)
		entity.Msg = enum.Accepted.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}

}

func (h *UserHandler) UserListHandler(c *gin.Context) {
	var q query.ListQuery
	entity := utils.BaseReturnBody()
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	list, err := h.UserSrv.List(&q)
	total, err := h.UserSrv.GetTotal(&q)

	if err != nil {
		panic(err)
	}
	if q.PageSize == 0 {
		q.PageSize = 5
	}
	ret := int(total % q.PageSize)
	ret2 := int(total / q.PageSize)
	totalPage := 0
	if ret == 0 {
		totalPage = ret2
	} else {
		totalPage = ret2 + 1
	}
	var newList []*dto.User
	for _, item := range list {
		r := h.GetUser(*item)
		newList = append(newList, &r)
	}

	entity = utils.ReturnBody(total,totalPage,enum.Success,newList)
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *UserHandler) AddUserHandler(c *gin.Context) {
	entity :=utils.BaseReturnBody()
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	r, err := h.UserSrv.Add(u)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.UserId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = int(enum.Success)
	entity.Msg = enum.Success.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}
