package models

type Project struct {
	HostURL     string `json:"host_url"`
	ID          int    `json:"id"`
	ImageURL    string `json:"image_url"`
	IsHighlight bool   `json:"is_highlight"`
	IsSleep     bool   `json:"is_sleep"`
	Language    string `json:"language"`
	Name        string `json:"name"`
	SourceURL   string `json:"source_url"`
	Status      string `json:"status"`
	Tools       string `json:"tools"`
}
