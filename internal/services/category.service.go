package services

import (
	"github.com/go-playground/validator/v10"
	"restaurant_management/internal/models/dto"
	"restaurant_management/internal/repositories"
)

type CategoryService interface {
	FindAll(request *dto.CategorySearch) ([]dto.CategoryResponse, error)
	FindById(categoryId string) (*dto.CategoryResponse, error)
	Create(request *dto.CategoryCreateRequest) (*dto.CategoryResponse, error)
	Edit(request *dto.CategoryUpdateRequest) (*dto.CategoryResponse, error)
	Delete(categoryId string) error
}

type CategoryServiceImpl struct {
	repository repositories.CategoryRepository
	validator  *validator.Validate
}

func (s CategoryServiceImpl) FindAll(request *dto.CategorySearch) ([]dto.CategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CategoryServiceImpl) FindById(categoryId string) (*dto.CategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CategoryServiceImpl) Create(request *dto.CategoryCreateRequest) (*dto.CategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CategoryServiceImpl) Edit(request *dto.CategoryUpdateRequest) (*dto.CategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CategoryServiceImpl) Delete(categoryId string) error {
	//TODO implement me
	panic("implement me")
}
