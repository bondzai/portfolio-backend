package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/bondzai/portfolio-backend/internal/utils"
	"github.com/valyala/fasthttp"
)

var apiUrl = utils.GetEnv("GO_WAKATIME_URL", "https://wakatime.com/api/v1/users/current/stats/all_time")
var apiKey = utils.GetEnv("GO_WAKATIME_API_KEY", "")

func cleanData(data []interface{}, newLastIndex int) []map[string]interface{} {
	cleanedData := make([]map[string]interface{}, newLastIndex+1)

	var totalPercent, totalHours, totalMinutes float64

	for i, item := range data {
		itemMap := item.(map[string]interface{})
		itemName := itemMap["name"].(string)
		itemPercent := itemMap["percent"].(float64)
		itemHours := itemMap["hours"].(float64)
		itemMinutes := itemMap["minutes"].(float64)
		itemText := itemMap["text"].(string)

		cleanedItem := map[string]interface{}{
			"name":    itemName,
			"percent": itemPercent,
			"hours":   itemHours,
			"minutes": itemMinutes,
			"text":    itemText,
		}

		if i < newLastIndex {
			totalPercent += itemPercent
			totalHours += itemHours
			totalMinutes += itemMinutes
		} else {
			otherPercent := 100.0 - totalPercent
			otherHours := totalHours * otherPercent / 100.0
			otherMinutes := totalMinutes * otherPercent / 100.0
			cleanedItem = map[string]interface{}{
				"name":    "Other",
				"percent": fmt.Sprintf("%.2f", otherPercent),
				"hours":   otherHours,
				"minutes": otherMinutes,
				"text":    fmt.Sprintf("%d hrs %d mins", int(otherHours), int(otherMinutes)),
			}
		}

		if i > newLastIndex {
			continue
		}
		cleanedData[i] = cleanedItem
	}

	return cleanedData
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
		"operating_systems":                             cleanData(response["data"].(map[string]interface{})["operating_systems"].([]interface{}), 2),
		"editors":                                       cleanData(response["data"].(map[string]interface{})["editors"].([]interface{}), 2),
		"languages":                                     cleanData(response["data"].(map[string]interface{})["languages"].([]interface{}), 3),
		"best_day":                                      response["data"].(map[string]interface{})["best_day"],
	}

	return wakatimeData, nil
}

var Wakatime, _ = FetchDataFromAPI()
