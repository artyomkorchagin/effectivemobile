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
	Host     string `mapstructure:"DB_HOST" validate:"required,hostname|ip"`
	Port     uint   `mapstructure:"DB_PORT" validate:"port,required"`
	User     string `mapstructure:"DB_USER" validate:"required"`
	Password string `mapstructure:"DB_PASSWORD" validate:"required"`
	Name     string `mapstructure:"DB_NAME" validate:"required"`
	SSLMode  string `mapstructure:"DB_SSLMODE" validate:"oneof=disable allow prefer require verify-ca verify-full,required"`
}

type ServerConfig struct {
	Host string `mapstructure:"SERVER_HOST" validate:"hostname|ip"`
	Port uint   `mapstructure:"SERVER_PORT" validate:"port,required"`
}

func LoadConfig(validate *validator.Validate) (*Config, error) {
	viper.SetConfigFile(".env")

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", 8080)
	viper.SetDefault("DB_SSLMODE", "disable")

	_ = viper.ReadInConfig()

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
