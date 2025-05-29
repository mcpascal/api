package logger

import (
	"api/configs"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.SugaredLogger
)

func Setup() {
	config := configs.App.LoggerInfo

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	writeSyncer := getWriter(config.Path)
	level := getLevel(config.Level)

	// 创建 core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// 创建 logger
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
}

func getLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func getWriter(filepath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	// 同时写入日志文件 + 控制台（开发环境）
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(lumberJackLogger),
		zapcore.AddSync(os.Stdout),
	)
}
