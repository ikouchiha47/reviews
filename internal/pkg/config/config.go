package config

import (
	"fmt"
	"reviews/pkg/db"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Address         string
	LogLevel        string
	LogFormat       string
	ShutDownTimeout time.Duration
	NewRelicKey     string
	NewRelicEnabled bool
	NewrelicAppName string
	DBConfig        db.Config
	TranslationPath string
}

func setDefaults() {
	viper.SetDefault("APP_NAME", "reviewsapp")
	viper.SetDefault("APP_PORT", "8000")
}

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	viper.AddConfigPath("./")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func New() *Config {
	return newConfig()
}

func newConfig() *Config {
	return &Config{
		Address:         ":9090",
		LogLevel:        "info",
		ShutDownTimeout: time.Second * 5,
		DBConfig:        NewDBConfig(),
		TranslationPath: "./reviews/asset/i18n/definations",
	}
}

func MigrationPath() string {
	return ""
}
