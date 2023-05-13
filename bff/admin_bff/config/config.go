package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HttpPort int `json:"http_port" mapstructure:"http_port"`

	BranchServiceHost    string `json:"branch_service_host" mapstructure:"branch_service_host"`
	BranchServicePort    int    `json:"branch_service_port" mapstructure:"branch_service_port"`
	AccountServiceHost   string `json:"account_service_host" mapstructure:"account_service_host"`
	AccountServicePort   int    `json:"account_service_port" mapstructure:"account_service_port"`
	StaffServiceHost     string `json:"staff_service_host" mapstructure:"staff_service_host"`
	StaffServicePort     int    `json:"staff_service_port" mapstructure:"staff_service_port"`
	GoodsServiceHost     string `json:"goods_service_host" mapstructure:"goods_service_host"`
	GoodsServicePort     int    `json:"goods_service_port" mapstructure:"goods_service_port"`
	EventServiceHost     string `json:"event_service_host" mapstructure:"event_service_host"`
	EventServicePort     int    `json:"event_service_port" mapstructure:"event_service_port"`
	WarehouseServiceHost string `json:"warehouse_service_host" mapstructure:"warehouse_service_host"`
	WarehouseServicePort int    `json:"warehouse_service_port" mapstructure:"warehouse_service_port"`
	OrderServiceHost     string `json:"order_service_host" mapstructure:"order_service_host"`
	OrderServicePort     int    `json:"order_service_port" mapstructure:"order_service_port"`
	StatisticServiceHost string `json:"statistic_service_host" mapstructure:"statistic_service_host"`
	StatisticServicePort int    `json:"statistic_service_port" mapstructure:"statistic_service_port"`
}

func Load() (config *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Admin BFF load default config")
		return loadDefaultConfig(), nil
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func loadDefaultConfig() *Config {
	return &Config{
		HttpPort:             10000,
		BranchServicePort:    14000,
		AccountServicePort:   14083,
		StaffServicePort:     14082,
		GoodsServicePort:     14080,
		EventServicePort:     14060,
		WarehouseServicePort: 14081,
		OrderServicePort:     14070,
		StatisticServicePort: 14090,
	}
}
