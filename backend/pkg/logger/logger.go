package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type ILogger interface {
	getLevel() log.Level
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
}

func (i ILogger) Fatalf(s string, err error) {
	panic("unimplemented")
}

var Logger ILogger

type LoggerConfig struct {
	LogLevel string `yaml:"log_level"`
}
type appLogger struct {
	level  string
	logger *log.Logger
}

var loggerLevelMap = map[string]log.Level{
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
	"panic": log.PanicLevel,
	"fatal": log.FatalLevel,
	"trace": log.TraceLevel,
}

func (l *appLogger) getLevel() log.Level {
	if level, exists := loggerLevelMap[l.level]; exists {
		return level
	}
	return log.InfoLevel
}

func InitLogger(config *LoggerConfig) ILogger {
	logger := &appLogger{
		level:  config.LogLevel,
		logger: log.New(),
	}
	logger.logger = log.StandardLogger()
	logLevel := logger.getLevel()

	env := os.Getenv("APP_ENV")

	if env == "production" {
		logger.logger.SetFormatter(&log.JSONFormatter{})
	} else {
		logger.logger.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
		})
	}
	logger.logger.SetLevel(logLevel)
	return logger
}

func (l *appLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}
func (l *appLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *appLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}
func (l *appLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *appLogger) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}
func (l *appLogger) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

func (l *appLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}
func (l *appLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *appLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}
func (l *appLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *appLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}
func (l *appLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *appLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}
func (l *appLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
