package utils

import (
	"github.com/korasdor/todo-app/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

/*
*tessfs
 */
func InitConfigs() error {
	viper.SetConfigFile(consts.ENV_FILE)
	viper.AddConfigPath(consts.ENV_FILE_DIRECTORY)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.AutomaticEnv()

	return nil
}

func GetEnvVar(name string) string {
	if !viper.IsSet(name) {
		logrus.Errorf("Environment variable %s is not set", name)
	}

	value := viper.GetString(name)
	return value
}
