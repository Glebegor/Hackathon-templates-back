package bootstrap

import "github.com/sirupsen/logrus"

type Logger struct {
	LoggerFile string
}

func NewLogger(env *Env) (*Logger, error) {

	logger := Logger{}

	return &logger, nil
}

func (l *Logger) ToLogFile(message string) {
	logrus.Info(message)
}
