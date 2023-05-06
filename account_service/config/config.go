package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Env          string        `json:"env" mapstructure:"env"`
	HttpPort     int           `json:"http_port" mapstructure:"http_port"`
	MySQL        *MySQLConfig  `json:"mysql" mapstructure:"mysql"`
	ServiceFlags *ServiceFlags `json:"service_flags" mapstructure:"service_flags"`

	StaffServicePort    int `json:"staff_service_port" mapstructure:"staff_service_port"`
	CustomerServicePort int `json:"customer_service_port" mapstructure:"customer_service_port"`
}

type MySQLConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

type ServiceFlags struct {
	IsEnableAsync bool `json:"is_enable_async" mapstructure:"is_enable_async"`
}

func Load() (config *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Account Service load default config")
		return loadDefaultConfig(), nil
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func loadDefaultConfig() *Config {
	return &Config{
		Env:      "docker",
		HttpPort: 14083,
		MySQL: &MySQLConfig{
			Host:     "mysql",
			Port:     3306,
			Username: "bpel",
			Password: "bpel",
			Database: "account_service",
		},
		ServiceFlags: &ServiceFlags{
			IsEnableAsync: true,
		},
		StaffServicePort:    14082,
		CustomerServicePort: 14050,
	}
}
