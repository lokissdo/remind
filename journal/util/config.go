package util

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBSource     string        `mapstructure:"postgres"`
	ServerAddr   string 	   `mapstructure:"server_address"`
	Environment  string        `mapstructure:"environment"`
	MigrationURL string        `mapstructure:"migration_url"`
	RabbitMQURL  string        `mapstructure:"rabbitmq_url"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("journal")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	fmt.Printf("Config file used: %s\n", viper.ConfigFileUsed())
	fmt.Printf("DB source is: %s\n", viper.Get("postgres"))

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return config, err
}
