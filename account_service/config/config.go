package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HttpPort     int           `json:"http_port" mapstructure:"http_port"`
	MySQL        *MySQLConfig  `json:"mysql" mapstructure:"mysql"`
	ServiceFlags *ServiceFlags `json:"service_flags" mapstructure:"service_flags"`

	StaffServiceHost    string `json:"staff_service_host" mapstructure:"staff_service_host"`
	StaffServicePort    int    `json:"staff_service_port" mapstructure:"staff_service_port"`
	CustomerServiceHost string `json:"customer_service_host" mapstructure:"customer_service_host"`
	CustomerServicePort int    `json:"customer_service_port" mapstructure:"customer_service_port"`
	KafkaHost           string `json:"kafka_host" mapstructure:"kafka_host"`
	KafkaPort           int    `json:"kafka_port" mapstructure:"kafka_port"`
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
