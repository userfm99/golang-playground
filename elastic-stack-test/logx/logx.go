package logx

import (
	"go.uber.org/zap"
	"encoding/json"
	"time"
)

var rawJson = []byte(`
{
	"level": "debug",
	"encoding": "json",
	"outputPaths": ["stdout", "/tmp/logs"],
	"errorOutputPaths": ["stderr"],
	"encoderConfig": {
	"messageKey": "message",
	"levelKey": "level",
	"levelEncoder": "lowercase"
	  }
}
`)

var cfg zap.Config
var logger *zap.Logger

func init () {
	if err := json.Unmarshal(rawJson, &cfg); err != nil {
		panic(err)
	}
	cfg.OutputPaths = []string{"stdout", "/home/fadly/logs"}
	mLogger, err := cfg.Build()
	logger = mLogger
	if err != nil {
		panic(err)
	}
}

func Info(message string) {
	defer logger.Sync()
	logger.Info(message, zap.String("url", "elk-log-test"), zap.String("time", time.Now().String()))
}

func Debug(message string) {
	defer logger.Sync()
	logger.Debug(message, zap.String("url", "add"), zap.String("time", time.Now().String()))
}

func Err(message string) {
	defer logger.Sync()
	logger.Error(message, zap.String("url", "add"), zap.String("time", time.Now().String()))
}