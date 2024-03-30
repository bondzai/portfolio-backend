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

// var Wakatime = map[string]interface{}{
// 	"best_day": map[string]interface{}{
// 		"date":          "2022-11-12",
// 		"text":          "15 hrs 11 mins",
// 		"total_seconds": 54673.290055,
// 	},
// 	"days_including_holidays": 606,
// 	"editors": []map[string]interface{}{
// 		{
// 			"decimal":       "2021.75",
// 			"digital":       "2021:45",
// 			"hours":         2021,
// 			"minutes":       45,
// 			"name":          "VS Code",
// 			"percent":       83.08,
// 			"text":          "2,021 hrs 45 mins",
// 			"total_seconds": 7278336.450685,
// 		},
// 		{
// 			"decimal":       "202.00",
// 			"digital":       "202:00",
// 			"hours":         202,
// 			"minutes":       0,
// 			"name":          "Zsh",
// 			"percent":       8.3,
// 			"text":          "202 hrs",
// 			"total_seconds": 727235.265727,
// 		},
// 		{
// 			"decimal":       "164.98",
// 			"digital":       "164:59",
// 			"hours":         164,
// 			"minutes":       59,
// 			"name":          "Browser",
// 			"percent":       6.78,
// 			"text":          "164 hrs 59 mins",
// 			"total_seconds": 593984.678767,
// 		},
// 		{
// 			"decimal":       "44.73",
// 			"digital":       "44:44",
// 			"hours":         44,
// 			"minutes":       44,
// 			"name":          "PyCharmCore",
// 			"percent":       1.84,
// 			"text":          "44 hrs 44 mins",
// 			"total_seconds": 161059.438042,
// 		},
// 	},
// 	"human_readable_range":                          "since Jul 19 2022",
// 	"human_readable_total_including_other_language": "2,433 hrs 30 mins",
// 	"languages": []map[string]interface{}{
// 		{
// 			"decimal":       "838.10",
// 			"digital":       "838:06",
// 			"hours":         838,
// 			"minutes":       6,
// 			"name":          "Python",
// 			"percent":       34.44,
// 			"text":          "838 hrs 6 mins",
// 			"total_seconds": 3017180.493675,
// 		},
// 		{
// 			"decimal":       "588.92",
// 			"digital":       "588:55",
// 			"hours":         588,
// 			"minutes":       55,
// 			"name":          "Go",
// 			"percent":       24.2,
// 			"text":          "588 hrs 55 mins",
// 			"total_seconds": 2120151.298502,
// 		},
// 		{
// 			"decimal":       "337.45",
// 			"digital":       "337:27",
// 			"hours":         337,
// 			"minutes":       27,
// 			"name":          "JavaScript",
// 			"percent":       13.87,
// 			"text":          "337 hrs 27 mins",
// 			"total_seconds": 1214877.817893,
// 		},
// 	},
// 	"operating_systems": []map[string]interface{}{
// 		{
// 			"decimal":       "2377.63",
// 			"digital":       "2377:38",
// 			"hours":         2377,
// 			"minutes":       38,
// 			"name":          "Linux",
// 			"percent":       97.7,
// 			"text":          "2,377 hrs 38 mins",
// 			"total_seconds": 8559511.901786,
// 		},
// 		{
// 			"decimal":       "54.75",
// 			"digital":       "54:45",
// 			"hours":         54,
// 			"minutes":       45,
// 			"name":          "Windows",
// 			"percent":       2.25,
// 			"text":          "54 hrs 45 mins",
// 			"total_seconds": 197120.629818,
// 		},
// 		{
// 			"decimal":       "1.10",
// 			"digital":       "1:06",
// 			"hours":         1,
// 			"minutes":       6,
// 			"name":          "Android",
// 			"percent":       0.05,
// 			"text":          "1 hr 6 mins",
// 			"total_seconds": 3983.301617,
// 		},
// 	},
// }
