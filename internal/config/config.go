package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DB      DBConfig     `mapstructure:",squash"`
	Server  ServerConfig `mapstructure:",squash"`
	LogMode string       `mapstructure:"LOG_MODE"`
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST" validate:"ip"`
	Port     int    `mapstructure:"DB_PORT" validate:"port,required"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	SSLMode  string `mapstructure:"DB_SSLMODE" validate:"oneof=enable disable,required"`
}

type ServerConfig struct {
	Host string `mapstructure:"SERVER_HOST" validate:"ip"`
	Port string `mapstructure:"SERVER_PORT" validate:"port,required"`
}

func LoadConfig(validate *validator.Validate) (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return &cfg, nil
}

func (cfg *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.SSLMode)
}
