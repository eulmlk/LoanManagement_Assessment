package bootstrap

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitEnv() error {
	viper.SetConfigFile(".env") // Specify the file to read

	// Automatically look for environment variables that match
	viper.AutomaticEnv()

	// Reading in the config file
	return viper.ReadInConfig()
}

func GetEnv(key string) (string, error) {
	if !viper.IsSet(key) {
		return "", fmt.Errorf("environment variable %s not found", key)
	}

	return viper.GetString(key), nil
}
