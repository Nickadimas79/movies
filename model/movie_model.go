package model

type Movie struct {
	ID          int     `json:"id`
	Title       string  `json:"title"`
	ReleaseYear int     `json:"releaseYear"`
	Score       float64 `json:"score"`
}
