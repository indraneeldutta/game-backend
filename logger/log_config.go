package logger

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

// SetUpLogging sets up the logger to be used in the application
func SetUpLogging() {
	encoderConfig := ecszap.ECSCompatibleEncoderConfig(zap.NewDevelopmentEncoderConfig())
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, os.Stdout, zap.DebugLevel)
	l := zap.New(ecszap.WrapCore(core), zap.AddCaller())

	// l, _ := zap.NewProduction()

	Log = l.Sugar()
}
