package service

import (
	"context"
	"github.com/goodliving/usercenter/model"
	"github.com/goodliving/usercenter/pkg/logging"
	"github.com/goodliving/usercenter/pkg/util"
)

type LoginService int

type LoginArgs struct {
	Username string
	Password string
}

type Reply struct {
	TraceID string `json:"trace_id"`
	Data interface{} `json:"data"`
	Success bool `json:"success"`
	Msg  string      `json:"msg"`
}

// Login验证用户密码是否正确，同时生成token
func (l LoginService) Login(ctx context.Context, args *LoginArgs, reply *Reply) error {
	logging.ZapLogger.Infow("登录信息", "请求参数", args)
	//reqMeta := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	//resMeta := ctx.Value(share.ResMetaDataKey).(map[string]string)

	reply.TraceID = "traceId"

	user, err := model.CheckAuth(args.Username, args.Password)
	if err != nil {
		reply.Data = nil
		reply.Success = false
		reply.Msg = "登录失败，请检查"
		return nil
	}

	token, expireTime, _ := util.GenerateToken(args.Username, args.Password)

	data := make(map[string]interface{})
	data["user_id"] = user.ID
	data["display_name"] = user.DisplayName
	data["token"] = token
	data["expire_time"] = expireTime

	reply.Data = data
	reply.Success = true
	reply.Msg = "登陆成功"
	return nil
}