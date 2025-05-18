package logger

import (
	"github.com/rs/zerolog"
	"os"
)


type Loggger struct {
	Console *zerolog.Logger
	Logger *zerolog.Logger
}

func NewLogger() (*Loggger,*os.File,error) {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil,nil,err
	}
	logger := zerolog.New(file).With().Timestamp().Logger()

	consoleLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	consoleLogger = consoleLogger.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	return &Loggger{
		Logger: &logger,
		Console: &consoleLogger,
	},file,nil
}