package config

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig loads a configuration file of a certain name by looking
// through the passed in paths
func LoadConfig(configPaths []string, configName string) {
	viper.SetConfigName(configName)
	for _, configPath := range configPaths {
		viper.AddConfigPath(configPath)
	}
	err := viper.ReadInConfig()
	if err != nil {
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("Cannot load config file [%s] from paths:\n", configName))
		for _, configPath := range configPaths {
			buffer.WriteString(fmt.Sprintf("%s\n", configPath))
		}
		buffer.WriteString(fmt.Sprintf("Error is:\n%s", err.Error()))
		panic(buffer.String())
	}
}
