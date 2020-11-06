package service

import (
	"context"
	"github.com/goodliving/usercenter/pkg/util"
)

type TokenService struct {
}

type TokenArgs struct {
	Token string
}

func (t TokenService) RefreshToken(ctx context.Context, args *TokenArgs, reply *Reply) error {
	token := args.Token
	if token == "" {
		reply.Msg = "token不能为空"
		reply.Success = false
		reply.Data = nil
		return nil
	}

	newToken, expireTime, err := util.RefreshToken(token)
	if err != nil {
		reply.Msg = err.Error()
		reply.Success = false
		reply.Data = nil
		return nil
	}

	data := make(map[string]interface{})
	data["token"] = newToken
	data["expire_time"] = expireTime
	reply.Data = data
	reply.Success = true
	reply.Msg = "更新成功"
	reply.TraceID = "asdfasdf"
	return nil
}

