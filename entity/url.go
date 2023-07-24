package entity

import (
	"time"
)

type Url struct {
	Id          string    `json:"id,omitempty"`
	Short       string    `json:"short,omitempty"`
	Long        string    `json:"long,omitempty"`
	Deeplink    string    `json:"deeplink"`
	Android     string    `json:"android"`
	Ios         string    `json:"ios"`
	DateCreated time.Time `json:"date_created,omitempty"`
}
