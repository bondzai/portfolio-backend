package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	CorsHeader string `mapstructure:"CORS_HEADERS"`
	CorsMethod string `mapstructure:"CORS_METHOD"`
	CorsOrigin string `mapstructure:"CORS_ORIGIN"`
	WakaApiKey string `mapstructure:"WAKATIME_API_KEY"`
	WakaUrl    string `mapstructure:"WAKATIME_URL"`
	MongoUrl   string `mapstructure:"MONGODB_URL"`
	MongoDB    string `mapstructure:"MONGODB_DB"`
	MongoCol   string `mapstructure:"MONGODB_COL"`
	DevToken   string `mapstructure:"DEV_TOKEN"`
	ExtraToken string `mapstructure:"EXTRA_TOKEN"`
}

func setDefaults() {
	viper.SetDefault("PORT", "10000")
	viper.SetDefault("CORS_HEADERS", "*")
	viper.SetDefault("CORS_METHOD", "*")
	viper.SetDefault("CORS_ORIGIN", "*")
	viper.SetDefault("WAKATIME_API_KEY", "")
	viper.SetDefault("WAKATIME_URL", "")
	viper.SetDefault("MONGODB_URL", "")
	viper.SetDefault("MONGODB_DB", "")
	viper.SetDefault("MONGODB_COL", "")
	viper.SetDefault("DEV_TOKEN", "")
	viper.SetDefault("EXTRA_TOKEN", "")
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		setDefaults()

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
	})

	return cfg
}
