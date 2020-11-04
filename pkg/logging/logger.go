package logging

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

var (
	ZapLogger *zap.SugaredLogger
)

func getLoggerLevel(lvl string) zapcore.Level {
	levelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}

	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf(t.Format("2006-01-02 15:04:05")))
}

// 配置日志输出文件
func Setup(appName string) {
	level := getLoggerLevel("debug")

	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "logs/" + appName + ".log",
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})

	encoder := zap.NewProductionEncoderConfig()

	encoder.EncodeTime = formatEncodeTime
	// 打开日志的颜色
	encoder.EncodeLevel = zapcore.LowercaseColorLevelEncoder

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	ZapLogger = l.Sugar()

}