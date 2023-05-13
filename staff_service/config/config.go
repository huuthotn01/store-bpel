package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HttpPort int          `json:"http_port" mapstructure:"http_port"`
	MySQL    *MySQLConfig `json:"mysql" mapstructure:"mysql"`

	AccountServiceHost string `json:"account_service_host" mapstructure:"account_service_host"`
	AccountServicePort int    `json:"account_service_port" mapstructure:"account_service_port"`
	KafkaHost          string `json:"kafka_host" mapstructure:"kafka_host"`
	KafkaPort          int    `json:"kafka_port" mapstructure:"kafka_port"`
}

type MySQLConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
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
		log.Println("Staff Service load default config")
		return loadDefaultConfig(), nil
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func loadDefaultConfig() *Config {
	return &Config{
		HttpPort:           14082,
		AccountServicePort: 14083,
		MySQL: &MySQLConfig{
			Host:     "mysql",
			Port:     3306,
			Username: "bpel",
			Password: "bpel",
			Database: "staff_service",
		},
	}
}
