package model

type EntityMovie struct {
	Title  string `json:"Title" validate:"required"`
	Year   string `json:"Year" validate:"required"`
	ImdbID string `json:"imdbID" validate:"required"`
	Type   string `json:"Type" validate:"required"`
	Poster string `json:"Poster" validate:"required"`
}
