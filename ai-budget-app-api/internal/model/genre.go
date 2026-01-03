package model

import (
	"github.com/google/uuid"
)

// ジャンル
type Genre struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateGenreRequest struct {
	Name string `json:"name" validate:"required,min=1"`
}

type RegisteredGenre struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateGenreResponse struct {
	Genres []RegisteredGenre `json:"genres"`
}

type GetGenresResponse struct {
	Genres []RegisteredGenre `json:"genres"`
}
