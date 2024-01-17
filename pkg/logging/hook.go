package logging

import (
	"github.com/sirupsen/logrus"
	"io"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (wh *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range wh.Writer {
		_, writeErr := w.Write([]byte(line))
		if err != nil {
			panic(writeErr)
		}
	}
	return err
}

func (wh *writerHook) Levels() []logrus.Level {
	return wh.LogLevels
}
