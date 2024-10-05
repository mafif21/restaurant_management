package converters

import (
	dto "restaurant_management/internal/models/dtos"
	"restaurant_management/internal/models/entities"
)

func ContactToResponse(category *entities.Category) *dto.CategoryResponse {
	return &dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
