package entity

type MetaData struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	Url         string `json:"url,omitempty"`
	Type        string `json:"type,omitempty"`
}
