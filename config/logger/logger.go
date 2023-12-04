package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	logger *zerolog.Logger
)

func init() {
	lg := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

	logger = &lg
}

func Info(message string) {
	logger.Info().Msg(message)
}

func Error(err error) {
	logger.Err(err)
}

func Debug(message string) {
	logger.Debug().Msg(message)
}
