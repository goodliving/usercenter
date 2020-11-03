package main

import (
	"fmt"
	"github.com/goodliving/usercenter/apollo"
	"github.com/goodliving/usercenter/config"
	"github.com/goodliving/usercenter/model"
	"github.com/goodliving/usercenter/service"
	"github.com/goodliving/usercenter/util/ip"
	"github.com/rcrowley/go-metrics"
	"github.com/shima-park/agollo"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

func init() {
	apollo.Setup()
}

func main() {

	host, ipErr := ip.ExternalIP()
	if ipErr != nil {
		log.Fatal(ipErr)
	}
	addr := fmt.Sprintf("%s:%s", host, agollo.Get("rpcx.port"))

	fmt.Println("addr: ", addr)

	consulAddr := agollo.Get("rpcx.consul.addr")

	mysqlInfo := config.GetMysqlInfo()
	mysqlInfo.Host = "rm-wz9y8z9i6j34gicr5mo.mysql.rds.aliyuncs.com:3306"
	model.Setup(mysqlInfo.User, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo. DbName)

	s := server.NewServer()

	addRegistryPlugin(s, "rpcx", addr, consulAddr)

	_ = s.RegisterName("usercenter", new(service.LoginService), "")

	err := s.Serve("tcp", addr)
	if err != nil {
		panic(err)
	}
}

func addRegistryPlugin(s *server.Server, basePath, addr, consulAddr string) {

	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		ConsulServers:  []string{consulAddr},
		BasePath:       basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}

	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}

	s.Plugins.Add(r)
}