package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	level := zapcore.DebugLevel

	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic("Не удалось открыть файл логов: " + err.Error())
	}
	fileWriter := zapcore.AddSync(logFile)

	core := zapcore.NewTee(zapcore.NewCore(encoder, fileWriter, level))

	Logger = zap.New(core, zap.AddCaller())
}
