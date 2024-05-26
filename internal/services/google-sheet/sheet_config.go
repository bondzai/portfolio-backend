package googlesheet

import "github.com/bondzai/portfolio-backend/utils"

func getGoogleSheetURL() string {
	return utils.GetEnv("GO_SHEET_URL", "")
}
