package main

import (
	"github.com/goodliving/functions"
	"github.com/goodliving/usercenter/model"
	"github.com/goodliving/usercenter/pkg/logging"
	"github.com/goodliving/usercenter/service"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

func init() {
	functions.SetupApollo()
}

func main() {

	// mysql启动配置
	mysqlInfo := functions.GetMysqlInfo()
	mysqlInfo.Host = "rm-wz9y8z9i6j34gicr5mo.mysql.rds.aliyuncs.com:3306"
	model.Setup(mysqlInfo.User, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo. DbName)

	// rpcx启动配置
	rpcxInfo := functions.GetRpcxInfo()
	logging.Setup("demo")

	traceP := &serverplugin.OpenTracingPlugin{}

	s := server.NewServer()
	s.Plugins.Add(traceP)

	functions.AddConsulRegistryPlugin(s, rpcxInfo.RpcxBasePath, rpcxInfo.ServiceAddr, rpcxInfo.ConsulAddr)
	_ = s.RegisterName("usercenter", new(service.LoginService), "")

	err := s.Serve("tcp", rpcxInfo.ServiceAddr)
	if err != nil {
		panic(err)
	}
}