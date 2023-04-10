package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
var logInitFail string = "{\"level\":\"error\",\"ts\":\"%v\",\"caller\":\"pkg/logger/logger.go\",\"msg\":\"Error while creating logger %v\"}"

// Initializig logger with functionality of writing logs to a file.
func NewLogger() *zap.SugaredLogger {
	loggerConfig := zap.NewProductionEncoderConfig()
	loggerConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(loggerConfig)
	timeStamp := time.Now().Format("02-01-2006")
	logFile, err := os.OpenFile(fmt.Sprintf("./logs/number-storage-service-%s.log", timeStamp),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		log.Printf(logInitFail,timeStamp, err.Error())
		
	}
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	consoleEncoder := zapcore.NewConsoleEncoder(loggerConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger.Sugar()
}