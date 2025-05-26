package logger

import (
	"container-manager/internal/middleware"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	entry *logrus.Entry
}

func NewLogrusLogger() middleware.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		log.SetOutput(os.Stdout)
		log.Warn("Failed to log to file, using default stderr")
	}

	log.SetLevel(logrus.InfoLevel)

	return &logrusLogger{
		entry: logrus.NewEntry(log),
	}
}

func (l *logrusLogger) Info(args ...any) {
	l.entry.Info(args...)
}

func (l *logrusLogger) Error(args ...any) {
	l.entry.Error(args...)
}

func (l *logrusLogger) WithFields(fields map[string]any) middleware.Logger {
	return &logrusLogger{
		entry: l.entry.WithFields(logrus.Fields(fields)),
	}
}
