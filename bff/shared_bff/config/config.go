package config

import "github.com/spf13/viper"

type Config struct {
	HttpPort int `json:"http_port" mapstructure:"http_port"`

	GoodsServicePort int `json:"goods_service_port" mapstructure:"goods_service_port"`
	EventServicePort int `json:"event_service_port" mapstructure:"event_service_port"`
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
