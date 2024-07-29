package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port int    `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	GRPC struct {
		UserHost    string `env-required:"true" yaml:"user_host" env:"GRPC_AUTH_HOST"`
		UserPort    int    `env-required:"true" yaml:"user_port" env:"GRPC_AUTH_PORT"`
		AuthHost    string `env-required:"true" yaml:"auth_host" env:"GRPC_USER_HOST"`
		AuthPort    int    `env-required:"true" yaml:"auth_port" env:"GRPC_USER_PORT"`
		JournalHost string `env-required:"true" yaml:"journal_host" env:"GRPC_JOURNAL_HOST"`
		JournalPort int    `env-required:"true" yaml:"journal_port" env:"GRPC_JOURNAL_PORT"`
		TodoHost    string `env-required:"true" yaml:"todo_host" env:"GRPC_TODO_HOST"`
		TodoPort    int    `env-required:"true" yaml:"todo_port" env:"GRPC_TODO_PORT"`
		AIHost      string `env-required:"true" yaml:"ai_host" env:"GRPC_AI_HOST"`
		AIPort      int    `env-required:"true" yaml:"ai_port" env:"GRPC_AI_PORT"`
	}

	Environment struct {
		Name string `env-required:"true" yaml:"name"   env:"NAME"`
	}

	Config struct {
		App         `yaml:"app"`
		HTTP        `yaml:"http"`
		GRPC        `yaml:"grpc"`
		Environment `yaml:"environment"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = cleanenv.ReadConfig(dir+"/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
