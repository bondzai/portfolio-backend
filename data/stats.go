package data

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"os"

	"github.com/bondzai/goez/toolbox"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

var devMode = true
var apiUrl = GetEnv("GO_WAKATIME_URL", "https://wakatime.com/api/v1/users/current/stats/all_time")
var apiKey = GetEnv("GO_WAKATIME_API_KEY", "")

func GetEnv(key, fallback string) string {
	if devMode {
		if err := godotenv.Load(); err != nil {
			log.Printf("Error loading .env file: %s\n", err)
		}
	}

	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func FetchDataFromAPI() (map[string]interface{}, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(apiUrl)
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(apiKey)))

	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, err
	}

	wakatimeData := map[string]interface{}{
		"human_readable_range":                          response["data"].(map[string]interface{})["human_readable_range"],
		"days_including_holidays":                       response["data"].(map[string]interface{})["days_including_holidays"],
		"human_readable_total_including_other_language": response["data"].(map[string]interface{})["human_readable_total_including_other_language"],
		"operating_systems":                             response["data"].(map[string]interface{})["operating_systems"],
		"editors":                                       response["data"].(map[string]interface{})["editors"],
		"languages":                                     response["data"].(map[string]interface{})["languages"],
		"best_day":                                      response["data"].(map[string]interface{})["best_day"],
	}

	toolbox.PPrint(wakatimeData)
	return wakatimeData, nil
}

var Wakatime, _ = FetchDataFromAPI()
