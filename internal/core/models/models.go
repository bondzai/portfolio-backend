package models

type Certification struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	OtherURL    string `json:"other_url"`
}

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

type Skill struct {
	ID        int    `json:"id"`
	ImageURL  string `json:"image_url"`
	IsShowing bool   `json:"is_showing"`
	Name      string `json:"name"`
	Topic     string `json:"topic"`
	URL       string `json:"url"`
}
