package config

import "github.com/spf13/viper"

type Config struct {
	HttpPort int `json:"http_port" mapstructure:"http_port"`

	CustomerServicePort int `json:"customer_service_port" mapstructure:"customer_service_port"`
	OrderServicePort    int `json:"order_service_port" mapstructure:"order_service_port"`
	CartServicePort     int `json:"cart_service_port" mapstructure:"cart_service_port"`
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
