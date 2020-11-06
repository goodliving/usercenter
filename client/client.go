package main

import (
	"context"
	"fmt"
	"github.com/goodliving/usercenter/service"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
	"log"
	"time"
)

func main()  {
	d := client.NewConsulDiscovery("rpcx", "usercenter", []string{"47.56.227.160:8500"}, nil)
	xClient := client.NewXClient("usercenter", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	defer xClient.Close()

	for {
		//reply := &example.Reply{}
		loginArgs := &service.LoginArgs{
			Username: "pzm",
			Password: "123456",
		}
		reply := &service.Reply{}

		ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, map[string]string{"traceId": "123123"})
		ctx = context.WithValue(ctx, share.ResMetaDataKey, make(map[string]string))
		err := xClient.Call(ctx, "Login", loginArgs, reply)

		if err != nil {
			log.Printf("failed to call: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		//tokenArgs := &service.TokenArgs{Token: "123123123"}
		//
		//tokenErr := xClient.Call(ctx, "RefreshToken", tokenArgs, reply)
		//if tokenErr != nil {
		//	log.Printf("failed to call: %v\n", tokenErr)
		//	time.Sleep(5 * time.Second)
		//	continue
		//}

		fmt.Println("resMeta: ", ctx.Value(share.ResMetaDataKey))
		time.Sleep(5 * time.Second)
	}
}
