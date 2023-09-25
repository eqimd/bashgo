package config

import (
	"github.com/sirupsen/logrus"
)

type config struct {
	IsDebug bool
	Logger  *logrus.Logger
}

var Config = &config{}

func InitConfig(
	isDebug bool,
) {
	Config.Logger = logrus.StandardLogger()
	Config.Logger.SetLevel(logrus.FatalLevel)

	Config.IsDebug = isDebug
	if isDebug {
		Config.Logger.SetLevel(logrus.DebugLevel)
	}
}
