package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		NoColor:    false,
	}

	Logger = zerolog.New(output).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(zerolog.InfoLevel)
}
