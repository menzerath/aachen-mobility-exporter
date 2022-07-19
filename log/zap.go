package log

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger     *zap.Logger
	loggerInit sync.Once
)

// NewLogger creates a new zap.Logger on the first call and just returns the existing one afterwards.
func NewLogger(options ...zap.Option) *zap.Logger {
	loggerInit.Do(func() {
		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		defaultOptions := []zap.Option{
			zap.WithCaller(false),
			zap.AddStacktrace(zapcore.FatalLevel),
		}

		newLogger, err := config.Build(append(defaultOptions, options...)...)
		if err != nil {
			panic(fmt.Sprintf("failed to create logger: %s", err))
		}

		logger = newLogger
	})
	return logger
}
