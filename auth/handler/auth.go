package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"grocery/auth/model/access"
	pb "grocery/auth/proto"
	"strconv"
)
var (
	accessService access.Service
)
type Auth struct{}

func (a Auth) MakeAccessToken(ctx context.Context, request *pb.Request, response *pb.Response) error {
	log.Info("[MakeAccessToken] 收到创建token请求")

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatUint(request.UserId, 10),
		Name: request.UserName,
	})
	if err != nil {
		response.Error=&pb.Error{
			Detail: err.Error(),
		}
		log.Errorf("[MakeAccessToken] token生成失败，err：%s", err)
		return err
	}
	response.Token=token
	return nil
}

func (a *Auth) DelUserAccessToken(ctx context.Context, request *pb.Request, response *pb.Response) error {
	log.Infof("[DelUserAccessToken] 清除用户token")
	err := accessService.DelUserAccessToken(request.Token)
	if err != nil {
		response.Error = &pb.Error{
			Detail: err.Error(),
		}

		log.Infof("[DelUserAccessToken] 清除用户token失败，err：%s", err)
		return err
	}

	return nil
}

func (s *Auth) GetCachedAccessToken(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Infof("[GetCachedAccessToken] 获取缓存的token，%d", req.UserId)
	token, err := accessService.GetCachedAccessToken(&access.Subject{
		ID: strconv.FormatInt(int64(req.UserId), 10),
	})
	if err != nil {
		rsp.Error = &pb.Error{
			Detail: err.Error(),
		}

		log.Errorf("[GetCachedAccessToken] 获取缓存的token失败，err：%s", err)
		return err
	}

	rsp.Token = token
	return nil
}


