package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	HttpPort int `json:"http_port" mapstructure:"http_port"`

	CustomerServiceHost string `json:"customer_service_host" mapstructure:"customer_service_host"`
	CustomerServicePort int    `json:"customer_service_port" mapstructure:"customer_service_port"`
	OrderServiceHost    string `json:"order_service_host" mapstructure:"order_service_host"`
	OrderServicePort    int    `json:"order_service_port" mapstructure:"order_service_port"`
	CartServiceHost     string `json:"cart_service_host" mapstructure:"cart_service_host"`
	CartServicePort     int    `json:"cart_service_port" mapstructure:"cart_service_port"`
	AccountServiceHost  string `json:"account_service_host" mapstructure:"account_service_host"`
	AccountServicePort  int    `json:"account_service_port" mapstructure:"account_service_port"`
}

func Load() (config *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Customer BFF load default config")
		return loadDefaultConfig(), nil
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func loadDefaultConfig() *Config {
	return &Config{
		HttpPort:            20000,
		CustomerServicePort: 14050,
		OrderServicePort:    14070,
		CartServicePort:     14061,
		AccountServicePort:  14083,
	}
}
