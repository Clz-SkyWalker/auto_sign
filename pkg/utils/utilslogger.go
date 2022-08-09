package utils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	sugarLogger *zap.SugaredLogger
)

func AddLogger(l Errno, fields ...zap.Field) {
	switch l.Level {
	case zapcore.DebugLevel, zapcore.InfoLevel:
		sugarLogger.Desugar().Log(l.Level, l.Error(), fields...)
	default:
		fields = append(fields, zap.Stack("stack"))
		sugarLogger.Desugar().Log(l.Level, l.Error(), fields...)
	}

}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCallerSkip(3))
	sugarLogger = logger.Sugar()
}

func DeferSync() {
	sugarLogger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 在zap中加入Lumberjack支持
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/auto_sign.log",
		MaxSize:    1,     // 以 MB 为单位
		MaxBackups: 5,     // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxAge:     30,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
