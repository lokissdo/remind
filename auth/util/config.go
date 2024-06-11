package util

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type ServerAddress struct {
	HTTP string `mapstructure:"http"`
	GRPC string `mapstructure:"grpc"`
}

type PasetoConfig struct {
	TokenSymmetricKey    string        `mapstructure:"token_symmetric_key"`
	AccessTokenDuration  time.Duration `mapstructure:"access_token_duration"`
	RefreshTokenDuration time.Duration `mapstructure:"refresh_token_duration"`
}

type EmailConfig struct {
	EmailSenderName     string `mapstructure:"email_sender_name"`
	EmailSenderAddress  string `mapstructure:"email_sender_address"`
	EmailSenderPassword string `mapstructure:"email_sender_password"`
}

type Config struct {
	DBSource     string        `mapstructure:"postgres"`
	RedisAddress string        `mapstructure:"redis.address"`
	ServerAddr   ServerAddress `mapstructure:"server_address"`
	PasetoConfig PasetoConfig  `mapstructure:"paseto"`
	EmailConfig  EmailConfig   `mapstructure:"email"`
	Environment  string        `mapstructure:"environment"`
	MigrationURL string        `mapstructure:"migration_url"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("auth")
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
