package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

type EchoLogger struct {
	z zerolog.Logger
}

func NewEchoLogger(z zerolog.Logger) *EchoLogger {
	return &EchoLogger{z: z}
}

func (l *EchoLogger) Output() io.Writer {
	return zerolog.ConsoleWriter{Out: os.Stdout}
}

func (l *EchoLogger) SetOutput(_ io.Writer) {
	// no-op since zerolog doesn't support dynamic output change directly
}

func (l *EchoLogger) Prefix() string {
	return ""
}

func (l *EchoLogger) SetPrefix(_ string) {
	// no-op
}

func (l *EchoLogger) Level() log.Lvl {
	switch zerolog.GlobalLevel() {
	case zerolog.DebugLevel:
		return log.DEBUG
	case zerolog.InfoLevel:
		return log.INFO
	case zerolog.WarnLevel:
		return log.WARN
	case zerolog.ErrorLevel:
		return log.ERROR
	case zerolog.FatalLevel:
		return log.OFF
	default:
		return log.INFO
	}
}

func (l *EchoLogger) SetLevel(level log.Lvl) {
	var zerologLevel zerolog.Level

	switch level {
	case log.DEBUG:
		zerologLevel = zerolog.DebugLevel
	case log.INFO:
		zerologLevel = zerolog.InfoLevel
	case log.WARN:
		zerologLevel = zerolog.WarnLevel
	case log.ERROR:
		zerologLevel = zerolog.ErrorLevel
	case log.OFF:
		zerologLevel = zerolog.Disabled
	default:
		zerologLevel = zerolog.InfoLevel
	}
	newLogger := Log.Level(zerologLevel)
	Log = newLogger

}

func (l *EchoLogger) Printj(j log.JSON) {
	l.z.Info().Fields(j).Msg("print")
}

func (l *EchoLogger) Debugj(j log.JSON) {
	l.z.Debug().Fields(j).Msg("debug")
}

func (l *EchoLogger) Infoj(j log.JSON) {
	l.z.Info().Fields(j).Msg("info")
}

func (l *EchoLogger) Warnj(j log.JSON) {
	l.z.Warn().Fields(j).Msg("warn")
}

func (l *EchoLogger) Errorj(j log.JSON) {
	l.z.Error().Fields(j).Msg("error")
}

func (l *EchoLogger) Fatalj(j log.JSON) {
	l.z.Fatal().Fields(j).Msg("fatal")
}

func (l *EchoLogger) Panicj(j log.JSON) {
	l.z.Panic().Fields(j).Msg("panic")
}

// SetHeader is a no-op as zerolog doesn't use headers in the same way
func (l *EchoLogger) SetHeader(_ string) {
	// no-op as zerolog uses a different logging format
}

func (l *EchoLogger) Print(i ...interface{}) {
	l.z.Info().Msg(fmt.Sprint(i...))
}

func (l *EchoLogger) Printf(format string, args ...interface{}) {
	l.z.Info().Msgf(format, args...)
}

func (l *EchoLogger) Debug(i ...interface{}) {
	l.z.Debug().Msg(fmt.Sprint(i...))
}

func (l *EchoLogger) Debugf(format string, args ...interface{}) {
	l.z.Debug().Msgf(format, args...)
}

func (l *EchoLogger) Info(i ...interface{}) {
	l.z.Info().Msg(fmt.Sprint(i...))
}

func (l *EchoLogger) Infof(format string, args ...interface{}) {
	l.z.Info().Msgf(format, args...)
}

func (l *EchoLogger) Warn(i ...interface{}) {
	l.z.Warn().Msg(fmt.Sprint(i...))
}

func (l *EchoLogger) Warnf(format string, args ...interface{}) {
	l.z.Warn().Msgf(format, args...)
}

func (l *EchoLogger) Error(i ...interface{}) {
	l.z.Error().Msg(fmt.Sprint(i...))
}

func (l *EchoLogger) Errorf(format string, args ...interface{}) {
	l.z.Error().Msgf(format, args...)
}

func (l *EchoLogger) Fatal(i ...interface{}) {
	l.z.Fatal().Msg(fmt.Sprint(i...))
}

func (l *EchoLogger) Fatalf(format string, args ...interface{}) {
	l.z.Fatal().Msgf(format, args...)
}

func (l *EchoLogger) Panic(i ...interface{}) {
	l.z.Panic().Msg(fmt.Sprint(i...))
}

func (l *EchoLogger) Panicf(format string, args ...interface{}) {
	l.z.Panic().Msgf(format, args...)
}
