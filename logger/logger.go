package logger

import (
	"io"
	"os"
	"time"

	"github.com/sainp7/go-mongo-echo-lab/config"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger(cfg *config.AppConfig) {
	var output io.Writer = os.Stdout

	if cfg.LogFormat == "console" {
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	}

	level, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = zerolog.InfoLevel // fallback
	}

	zerolog.SetGlobalLevel(level)

	Log = zerolog.New(output).Level(level).With().Timestamp().Logger()
}
