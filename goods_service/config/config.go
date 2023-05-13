package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HttpPort int          `json:"http_port" mapstructure:"http_port"`
	MySQL    *MySQLConfig `json:"mysql" mapstructure:"mysql"`

	WarehouseServiceHost string `json:"warehouse_service_host" mapstructure:"warehouse_service_host"`
	WarehouseServicePort int    `json:"warehouse_service_port" mapstructure:"warehouse_service_port"`
	EventServiceHost     string `json:"event_service_host" mapstructure:"event_service_host"`
	EventServicePort     int    `json:"event_service_port" mapstructure:"event_service_port"`
	OrderServiceHost     string `json:"order_service_host" mapstructure:"order_service_host"`
	OrderServicePort     int    `json:"order_service_port" mapstructure:"order_service_port"`
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
		log.Println("Goods Service load default config")
		return loadDefaultConfig(), nil
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func loadDefaultConfig() *Config {
	return &Config{
		HttpPort:             14080,
		WarehouseServicePort: 14081,
		EventServicePort:     14060,
		OrderServicePort:     14070,
		MySQL: &MySQLConfig{
			Host:     "mysql",
			Port:     3306,
			Username: "bpel",
			Password: "bpel",
			Database: "goods_service",
		},
	}
}
