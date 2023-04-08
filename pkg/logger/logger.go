package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger() *Logger {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{logger: logger}
}

func (logger *Logger) Error() *zerolog.Event {
	return logger.logger.Error()
}

func (logger *Logger) Info() *zerolog.Event {
	return logger.logger.Info()
}

func (logger *Logger) Debug() *zerolog.Event {
	return logger.logger.Debug()
}

func (logger *Logger) Warn() *zerolog.Event {
	return logger.logger.Warn()
}

func (logger *Logger) Fatal() *zerolog.Event {
	return logger.logger.Fatal()
}

func (logger *Logger) Panic() *zerolog.Event {
	return logger.logger.Panic()
}
