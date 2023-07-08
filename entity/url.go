package entity

import (
	"time"
)

type Url struct {
	Id          string    `json:"id,omitempty"`
	Short       string    `json:"short,omitempty"`
	Long        string    `json:"long,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
}
