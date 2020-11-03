package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goodliving/usercenter/model"
	"github.com/smallnest/rpcx/share"
)

type LoginService int

type LoginArgs struct {
	Username string
	Password string
}

type LoginReply struct {
	Code int         `json:"code"`
	TraceID string `json:"trace_id"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Success bool `json:"success"`
}

func (l LoginService) Login(ctx context.Context, args *LoginArgs, reply *LoginReply) error {
	fmt.Println("登录信息： ", args)
	reqMeta := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	resMeta := ctx.Value(share.ResMetaDataKey).(map[string]string)

	fmt.Printf("received meta: %+v\n", reqMeta["traceId"])
	resMeta["echo"] = "from server"

	user, err := model.CheckAuth(args.Username, args.Password)

	u, _ := json.Marshal(user)
	if err != nil {
		reply.Data = nil
		reply.Success = false
		return nil
	}

	reply.Data = string(u)
	reply.Success = true
	return nil
}