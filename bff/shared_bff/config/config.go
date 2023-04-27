package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HttpPort int `json:"http_port" mapstructure:"http_port"`

	GoodsServicePort   int `json:"goods_service_port" mapstructure:"goods_service_port"`
	EventServicePort   int `json:"event_service_port" mapstructure:"event_service_port"`
	AccountServicePort int `json:"account_service_port" mapstructure:"account_service_port"`
	OrderServicePort   int `json:"order_service_port" mapstructure:"order_service_port"`
}

func Load() (config *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Shared BFF load default config")
		return loadDefaultConfig(), nil
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func loadDefaultConfig() *Config {
	return &Config{
		HttpPort:           30000,
		GoodsServicePort:   14080,
		EventServicePort:   14060,
		AccountServicePort: 14083,
		OrderServicePort:   14070,
	}
}
