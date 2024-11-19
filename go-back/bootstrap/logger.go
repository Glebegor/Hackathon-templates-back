package bootstrap

import "github.com/sirupsen/logrus"

type Logger struct {
	LoggerFile string
}

func NewLogger(env *Env) (*Logger, error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logger := Logger{}

	return &logger, nil
}
