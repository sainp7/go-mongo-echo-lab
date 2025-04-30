package logger

import (
	"github.com/rs/zerolog"
)

type MongoDBLogger struct {
	logger zerolog.Logger
}

func NewMongoDBLogger(z zerolog.Logger) *MongoDBLogger {
	return &MongoDBLogger{logger: z}
}

func (m *MongoDBLogger) Info(level int, message string, keysAndValues ...interface{}) {
	switch level {
	case 1:
		m.info(message, keysAndValues...)
	case 2:
		m.debug(message, keysAndValues...)
	default:
		return
	}
}

func (m *MongoDBLogger) Error(err error, message string, keysAndValues ...interface{}) {
	m.logger.Error().Err(err).Fields(makeFields(keysAndValues...)).Msg(message)
}

func (m *MongoDBLogger) info(msg string, keysAndValues ...interface{}) {
	m.logger.Info().Fields(makeFields(keysAndValues...)).Msg(msg)
}

func (m *MongoDBLogger) debug(msg string, keysAndValues ...interface{}) {
	m.logger.Debug().Fields(makeFields(keysAndValues...)).Msg(msg)
}

// makeFields converts key-value pairs into a map for zerolog
func makeFields(keysAndValues ...interface{}) map[string]interface{} {
	fields := make(map[string]interface{})
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			fields[keysAndValues[i].(string)] = keysAndValues[i+1]
		}
	}
	return fields
}
