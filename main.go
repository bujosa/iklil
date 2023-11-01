package main

import (
	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	var wrapper WraperLogger
	var l retryablehttp.LeveledLogger = &wrapper
	wrapper.sugar = sugar
	client := retryablehttp.NewClient()
	client.Logger = l
	client.Get("http://www.google.com")
}


// This is the interface that retryablehttp expects for logging purposes
type WraperLogger struct {
	sugar *zap.SugaredLogger
}

func (w WraperLogger) Error(msg string, keyAndValues ...interface{}) {
   w.sugar.Error(msg, keyAndValues)
}

func (w WraperLogger) Info(msg string, keyAndValues ...interface{}) {
   w.sugar.Info(msg, keyAndValues)
}

func (w WraperLogger) Debug(msg string, keyAndValues ...interface{}) {
   w.sugar.Debug(msg, keyAndValues)
}

func (w WraperLogger) Warn(msg string, keyAndValues ...interface{} ) {
   w.sugar.Warn(msg, keyAndValues)
}