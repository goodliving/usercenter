package auth_service

import (
	"context"
	"fmt"
	"github.com/goodliving/usercenter/model"
)

type AuthService struct {

}

type LoginArgs struct {
	Username string
	Password string
}

type Reply struct {
	User *model.Users
}

type Login int

func (l Login) Login(ctx context.Context, args *LoginArgs, reply *Reply) error {
	fmt.Println("登录信息： ", args)
	user, err := model.CheckAuth(args.Username, args.Password)
	if err != nil {
		reply.User = nil
		return nil
	}

	reply.User = user
	return nil
}