package logging

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"time"
)

type logger struct {
	*logrus.Logger
}

func GetLogger(ctx context.Context) Logger {
	return loggerFromContext(ctx)
}

func NewLogger() Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetReportCaller(true)
	logrusLogger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("../../logs", 0644)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("../../logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	logrusLogger.SetOutput(io.Discard)
	logrusLogger.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})
	logrusLogger.SetLevel(logrus.TraceLevel)

	return &logger{
		Logger: logrusLogger,
	}
}

func (l *logger) SetLevel(level logrus.Level) {
	l.Logger.SetLevel(level)
}

func (l *logger) GetLevel() logrus.Level {
	return l.Logger.GetLevel()
}

func (l *logger) WithField(key string, value interface{}) *logrus.Entry {
	return l.Logger.WithField(key, value)
}

func (l *logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.Logger.WithFields(fields)
}

func (l *logger) WithError(err error) *logrus.Entry {
	return l.Logger.WithError(err)
}

func (l *logger) WithContext(ctx context.Context) *logrus.Entry {
	return l.Logger.WithContext(ctx)
}

func (l *logger) WithTime(t time.Time) *logrus.Entry {
	return l.Logger.WithTime(t)
}

func (l *logger) Tracef(format string, args ...interface{}) {
	l.Logger.Tracef(format, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *logger) Warningf(format string, args ...interface{}) {
	l.Logger.Warningf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.Logger.Panicf(format, args...)
}

func (l *logger) Trace(args ...interface{}) {
	l.Logger.Traceln(args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.Logger.Debugln(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.Logger.Infoln(args...)
}

func (l *logger) Print(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l *logger) Warning(args ...interface{}) {
	l.Logger.Warningln(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.Logger.Errorln(args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.Logger.Fatalln(args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.Logger.Panicln(args...)
}
