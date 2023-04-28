package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HttpPort int          `json:"http_port" mapstructure:"http_port"`
	MySQL    *MySQLConfig `json:"mysql" mapstructure:"mysql"`

	GoodsServicePort int `json:"goods_service_port" mapstructure:"goods_service_port"`
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
		log.Println("Cart Service load default config")
		return loadDefaultConfig(), nil
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func loadDefaultConfig() *Config {
	return &Config{
		HttpPort:         14061,
		GoodsServicePort: 14080,
		MySQL: &MySQLConfig{
			Host:     "mysql",
			Port:     3306,
			Username: "bpel",
			Password: "bpel",
			Database: "cart_service",
		},
	}
}
