module github.com/goodliving/usercenter

go 1.14

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/goodliving/functions v0.0.0-20201104054351-f501404352e0
	github.com/jinzhu/gorm v1.9.16
	github.com/segmentio/ksuid v1.0.3
	github.com/shima-park/agollo v1.2.7
	github.com/smallnest/rpcx v0.0.0-20201027145221-c31b15be63d4
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0
