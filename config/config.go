package config

import (
	"flag"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	CorsHeader string
	CorsMethod string
	CorsOrigin string
	WakaApiKey string
	WakaUrl    string
	MongoUrl   string
	MongoDB    string
	MongoCol   string
	DevToken   string
	ExtraToken string
}

func NewConfig() *Config {
	devFlag := flag.Bool("dev", false, "Start Dev flag")
	flag.Parse()

	if *devFlag {
		viper.SetConfigFile(".env")
	} else {
		viper.AutomaticEnv()
	}

	viper.ReadInConfig()

	return &Config{
		Port:       viper.GetString("GO_PORT"),
		CorsHeader: viper.GetString("GO_CORS_HEADERS"),
		CorsMethod: viper.GetString("GO_CORS_METHOD"),
		CorsOrigin: viper.GetString("GO_CORS_ORIGIN"),
		WakaApiKey: viper.GetString("GO_WAKATIME_API_KEY"),
		WakaUrl:    viper.GetString("GO_WAKATIME_API_KEY"),
		MongoUrl:   viper.GetString("GO_MONGODB_URL"),
		MongoDB:    viper.GetString("GO_MONGODB_DB"),
		MongoCol:   viper.GetString("GO_MONGODB_COL"),
		DevToken:   viper.GetString("GO_DEV_TOKEN"),
		ExtraToken: viper.GetString("GO_EXTRA_TOKEN"),
	}
}

var conf = NewConfig()

func GetConfig() *Config {
	return conf
}
