package main

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx-usercenter/service/auth_service"
	"time"
)

func main()  {
	d := client.NewConsulDiscovery("/rpcx", "Login", []string{"47.56.227.160:8500"}, nil)
	xClient := client.NewXClient("Login", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	defer xClient.Close()

	for {
		//reply := &example.Reply{}
		args := &auth_service.LoginArgs{
			Username: "pzm",
			Password: "123456",
		}
		reply := &auth_service.Reply{}

		err := xClient.Call(context.Background(), "Login", args, reply)
		if err != nil {
			log.Printf("failed to call: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Println("结果: ", reply.User.ID)
		time.Sleep(5 * time.Second)
	}
}
