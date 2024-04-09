package logger

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

// Logger is a logger :)
type Logger struct {
	*zerolog.Logger
}

var (
	logger Logger
	once   sync.Once
)

func Get() *Logger {

	once.Do(func() {
		err := godotenv.Load()

		file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to open log file")
		}
		multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, file)

		zeroLogger := zerolog.New(multi).With().Timestamp().Logger()
		globalStatus := os.Getenv("LOG_LEVEL")
		fmt.Println(globalStatus, "*************")

		switch globalStatus {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn", "warning":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "err", "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel) // log info and above by default
		}
		logger = Logger{&zeroLogger}
	})
	return &logger
}
