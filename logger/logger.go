package logger

import (
	"jachow/code1024/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(logConfig *config.LogConfig) {
	// TODO: 初始化日志
	encoder := getEncoder()
	writeSyner := getWriteSyncer(logConfig.Filename, logConfig.MaxSize, logConfig.MaxBackups, logConfig.MaxAge)	
	core := zapcore.NewCore(
		encoder,
		writeSyner,
		zapcore.DebugLevel,
	)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.TimeKey = "time"
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeDuration = zapcore.SecondsDurationEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	config.CallerKey = "caller"
	return zapcore.NewConsoleEncoder(config)
}

func getWriteSyncer(filename string, maxSize int, maxBackups int, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,   // 日志文件的位置
		MaxSize:    maxSize,    // 在进行 切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: maxBackups, // 保留旧文件的最大个数
		MaxAge:     maxAge,     // 保留旧文件的最大天数
		LocalTime:  true,       // 使用本地时间进行日志轮转，避免权限问题
		Compress:   false,      // 不压缩旧日志文件
	}
	return zapcore.AddSync(lumberJackLogger)
}