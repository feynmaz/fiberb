package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(debug bool) {
	var cfg zap.Config

	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	if debug {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	cfg.Encoding = "json"
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}

	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)
}
