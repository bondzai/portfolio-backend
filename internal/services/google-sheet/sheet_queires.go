package googlesheet

import (
	"encoding/json"
	"net/http"

	"github.com/bondzai/portfolio-backend/internal/models"
)

func GetDataFromAPI(dataType string) (models.KeyValueSlice, error) {
	url := getGoogleSheetURL() + dataType

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data models.KeyValueSlice
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	data.Filter("is_showing", false)

	return data, nil
}
