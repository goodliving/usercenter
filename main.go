package main

import (
	"github.com/goodliving/functions"
	"github.com/goodliving/usercenter/model"
	"github.com/goodliving/usercenter/service"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

func init() {
	functions.SetupApollo()
}

func main() {

	rpcxAddr := functions.GetRpcxAddr()

	mysqlInfo := functions.GetMysqlInfo()
	mysqlInfo.Host = "rm-wz9y8z9i6j34gicr5mo.mysql.rds.aliyuncs.com:3306"
	model.Setup(mysqlInfo.User, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo. DbName)

	s := server.NewServer()

	addRegistryPlugin(s, "rpcx", rpcxAddr.ServiceAddr, rpcxAddr.ConsulAddr)

	_ = s.RegisterName("usercenter", new(service.LoginService), "")

	err := s.Serve("tcp", rpcxAddr.ServiceAddr)
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