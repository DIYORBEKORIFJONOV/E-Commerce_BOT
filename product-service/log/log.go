package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	fileLogger    zerolog.Logger
	consoleLogger zerolog.Logger
}

type LogContext struct {
	fileLogger    zerolog.Logger
	consoleLogger zerolog.Logger
}

type Config struct {
	LogFilePath string 
	ServiceName string 
}

func New(config Config) (*Logger, *os.File, error) {
	file, err := os.OpenFile(config.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}

	baseFileLogger := zerolog.New(file).With().Timestamp().Str("service", config.ServiceName).Logger()
	baseConsoleLogger := zerolog.New(os.Stdout).With().Timestamp().Str("service", config.ServiceName).Logger()
	baseConsoleLogger = baseConsoleLogger.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	return &Logger{
		fileLogger:    baseFileLogger,
		consoleLogger: baseConsoleLogger,
	}, file, nil
}

func (l *Logger) WithContext(fields map[string]string) *LogContext {
	builder := l.fileLogger.With()
	builderConsole := l.consoleLogger.With()

	for k, v := range fields {
		builder = builder.Str(k, v)
		builderConsole = builderConsole.Str(k, v)
	}

	return &LogContext{
		fileLogger:    builder.Logger(),
		consoleLogger: builderConsole.Logger(),
	}
}

func (lc *LogContext) Info(msg string) {
	lc.consoleLogger.Info().Msg(msg)
	lc.fileLogger.Info().Msg(msg)
}

func (lc *LogContext) Warn(msg string) {
	lc.consoleLogger.Warn().Msg(msg)
	lc.fileLogger.Warn().Msg(msg)
}

func (lc *LogContext) Error(err error, msg string) {
	lc.consoleLogger.Err(err).Msg(msg)
	lc.fileLogger.Err(err).Msg(msg)
}
