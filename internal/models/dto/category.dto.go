package dto

import "time"

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryUpdateRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type CategorySearch struct {
	filters map[string]any
	page    int
}

type CategoryResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
