package models

type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	//intentionally using string as time , as data from json was not in go standard RFC
	PublicationDate string   `json:"publication_date"`
	SourceName      string   `json:"source_name"`
	Category        []string `json:"category"`
	RelevanceScore  float64  `json:"relevance_score"`
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
	Entities        []string `json:"entities"`
}
