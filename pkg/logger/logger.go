package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	once      sync.Once
	logger    zerolog.Logger
	logFilter zerolog.HookFunc
)

func Logger() *zerolog.Logger {
	return &logger
}

func Info() *zerolog.Event {
	return Logger().Info()
}

func Error() *zerolog.Event {
	return Logger().Error()
}

func Fatal() *zerolog.Event {
	return Logger().Fatal()
}

func Debug() *zerolog.Event {
	return Logger().Debug()
}

func Panic() *zerolog.Event {
	return Logger().Panic()
}

func Trace() *zerolog.Event {
	return Logger().Trace()
}

func Warn() *zerolog.Event {
	return Logger().Warn()
}

func init() {
	once.Do(func() {
		writer := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("drop logs %d", missed)
		})

		env := os.Getenv("ENJOT_ENVIRONMENT")

		logLevel := os.Getenv("ENJOT_LOGGER_LEVEL")

		if env == "local" {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		} else if logLevel != "" {
			level, err := zerolog.ParseLevel(logLevel)

			if err != nil {
				zerolog.SetGlobalLevel(zerolog.InfoLevel)
			} else {
				zerolog.SetGlobalLevel(level)
			}

		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		logger = zerolog.New(writer).
			With().
			Timestamp().
			Logger()

	})
}
