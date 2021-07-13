package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

// SetUpLogging sets up the logger to be used in the application
func SetUpLogging() {
	l, _ := zap.NewProduction()

	Log = l.Sugar()
}
