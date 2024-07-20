package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var Env = LoadConfig()

type config struct {
	Port        string
	CorsHeader  string
	CorsMethod  string
	CorsOrigin  string
	WakaApiKey  string
	WakaUrl     string
	MongoUrl    string
	MongoDB     string
	DevToken    string
	ExtraToken  string
	RedisUrl    string
	RedisDb     int
	RedisUser   string
	RedisPass   string
	KafkaBroker string
	KafkaUser   string
	KafkaPass   string
	KafKaTopic  string
}

func LoadConfig() *config {
	var once sync.Once

	once.Do(func() {
		viper.SetConfigFile(".env")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading Config file, %s", err)
		}
	})

	return &config{
		Port:        viper.GetString("PORT"),
		CorsHeader:  viper.GetString("CORS_HEADERS"),
		CorsMethod:  viper.GetString("CORS_METHOD"),
		CorsOrigin:  viper.GetString("CORS_ORIGIN"),
		WakaApiKey:  viper.GetString("WAKATIME_API_KEY"),
		WakaUrl:     viper.GetString("WAKATIME_URL"),
		MongoUrl:    viper.GetString("MONGODB_URL"),
		MongoDB:     viper.GetString("MONGODB_DB"),
		DevToken:    viper.GetString("DEV_TOKEN"),
		ExtraToken:  viper.GetString("EXTRA_TOKEN"),
		RedisUrl:    viper.GetString("REDIS_URL"),
		RedisDb:     viper.GetInt("REDIS_DB"),
		RedisUser:   viper.GetString("REDIS_USER"),
		RedisPass:   viper.GetString("REDIS_PASS"),
		KafkaBroker: viper.GetString("KAFKA_BROKER"),
		KafkaUser:   viper.GetString("KAFKA_USER"),
		KafkaPass:   viper.GetString("KAFKA_PASS"),
		KafKaTopic:  "uzhfeczb-default",
	}
}
