package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Certification struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Description string             `json:"description"`
	ImageURL    string             `json:"image_url"`
	Name        string             `json:"name"`
	OtherURL    string             `json:"other_url"`
}

type Project struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	HostURL     string             `json:"host_url"`
	ImageURL    string             `json:"image_url"`
	IsHighlight bool               `json:"is_highlight"`
	IsSleep     bool               `json:"is_sleep"`
	Language    string             `json:"language"`
	Name        string             `json:"name"`
	SourceURL   string             `json:"source_url"`
	Status      string             `json:"status"`
	Tools       string             `json:"tools"`
}

type Skill struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	ImageURL  string             `json:"image_url"`
	IsShowing bool               `json:"is_showing"`
	Name      string             `json:"name"`
	Topic     string             `json:"topic"`
	URL       string             `json:"url"`
}

type TotalUsers struct {
	Time       time.Time `json:"time"`
	TotalUsers int       `json:"total_users"`
}
