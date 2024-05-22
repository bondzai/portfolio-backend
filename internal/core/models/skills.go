package models

type Skill struct {
	ID        int    `json:"id"`
	ImageURL  string `json:"image_url"`
	IsShowing bool   `json:"is_showing"`
	Name      string `json:"name"`
	Topic     string `json:"topic"`
	URL       string `json:"url"`
}
