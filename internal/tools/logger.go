package tools

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	debug   func(v ...interface{})
	info    func(v ...interface{})
	warning func(v ...interface{})
	error   func(v ...interface{})
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New()

	return &Logger{
		debug:   logger.WithField("prefix", prefix).Debug,
		info:    logger.WithField("prefix", prefix).Info,
		warning: logger.WithField("prefix", prefix).Warning,
		error:   logger.WithField("prefix", prefix).Error,
		writer:  writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	log.Debug(v...)
}
func (l *Logger) Info(v ...interface{}) {
	log.Info(v...)
}
func (l *Logger) Warn(v ...interface{}) {
	log.Warning(v...)
}
func (l *Logger) Error(v ...interface{}) {
	log.Error(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	log.Warningf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)

	return logger
}
