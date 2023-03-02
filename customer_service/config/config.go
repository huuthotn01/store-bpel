package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort int `json:"http_port" mapstructure:"http_port"`
	MySQL *MySQLConfig `json:"mysql" mapstructure:"mysql"`

	StaffServicePort int `json:"staff_service_port" mapstructure:"staff_service_port"`
}

type MySQLConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int `json:"port" mapstructure:"port"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

func Load() (config *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
