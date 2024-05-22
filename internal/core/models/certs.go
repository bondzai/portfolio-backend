package models

type Certification struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	OtherURL    string `json:"other_url"`
}
