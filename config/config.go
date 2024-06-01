package config

import (
	"log"
	"reflect"
	"sync"

	"github.com/spf13/viper"
)

type config struct {
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

var (
	cfg  *config
	once sync.Once
)

func LoadConfig() *config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()

		cfg = &config{}
		setDefaults(cfg)

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(cfg); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
	})

	return cfg
}

func setDefaults(cfg *config) {
	elem := reflect.TypeOf(cfg).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		defaultValue := field.Tag.Get("default")

		if defaultValue != "" {
			viper.SetDefault(field.Tag.Get("mapstructure"), defaultValue)
		}
	}
}
