package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	log logger
)

type bookstoreLogger interface {
	//Interface Implementation for elasticsearch
	Printf(format string, v ...interface{})
	//Interface Implementation for MySql
	Print(v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {

	logConfig := zap.Config{
		OutputPaths: []string{getOutputPath()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		Encoding: "json",
	}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogger() bookstoreLogger {
	return log
}

func getLevel() zapcore.Level {

	switch strings.TrimSpace(strings.ToLower(os.Getenv(envLogLevel))) {
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "debug":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}

}

func getOutputPath() string {

	outputPath := strings.TrimSpace(strings.ToLower(os.Getenv(envLogOutput)))
	if outputPath == "" {
		return "stdout"
	}
	return outputPath

}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v...))
	}
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {

	tags = append(tags, zap.NamedError("error", err))

	log.log.Error(msg)
	log.log.Sync()
}
