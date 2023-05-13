package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HttpPort int `json:"http_port" mapstructure:"http_port"`

	GoodsServiceHost   string `json:"goods_service_host" mapstructure:"goods_service_host"`
	GoodsServicePort   int    `json:"goods_service_port" mapstructure:"goods_service_port"`
	EventServiceHost   string `json:"event_service_host" mapstructure:"event_service_host"`
	EventServicePort   int    `json:"event_service_port" mapstructure:"event_service_port"`
	AccountServiceHost string `json:"account_service_host" mapstructure:"account_service_host"`
	AccountServicePort int    `json:"account_service_port" mapstructure:"account_service_port"`
	OrderServiceHost   string `json:"order_service_host" mapstructure:"order_service_host"`
	OrderServicePort   int    `json:"order_service_port" mapstructure:"order_service_port"`
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
