package config

import "github.com/spf13/viper"

type Config struct {
	HttpPort int `json:"http_port" mapstructure:"http_port"`

	BranchServicePort    int `json:"branch_service_port" mapstructure:"branch_service_port"`
	AccountServicePort   int `json:"account_service_port" mapstructure:"account_service_port"`
	StaffServicePort     int `json:"staff_service_port" mapstructure:"staff_service_port"`
	GoodsServicePort     int `json:"goods_service_port" mapstructure:"goods_service_port"`
	EventServicePort     int `json:"event_service_port" mapstructure:"event_service_port"`
	WarehouseServicePort int `json:"warehouse_service_port" mapstructure:"warehouse_service_port"`
	OrderServicePort     int `json:"order_service_port" mapstructure:"order_service_port"`
	StatisticServicePort int `json:"statistic_service_port" mapstructure:"statistic_service_port"`
}

func Load() (config *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
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
