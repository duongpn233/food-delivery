package config

import (
	"os"

	"github.com/spf13/viper"
)

type SqlConfig struct {
	MysqlConnection string `mapstructure:"sqlconnection"`
}

type appConfig struct {
	SqlConfig SqlConfig `mapstructure:"mysql"`
}

func (c *appConfig) GetSqlConnection() string {
	return c.SqlConfig.MysqlConnection
}

func LoadConfig(path string) (*appConfig, error) {
	var config appConfig
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(path)
	viper.SetConfigType("toml")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.GetViper().Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
