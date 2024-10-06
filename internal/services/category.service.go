package services

import (
	"github.com/go-playground/validator/v10"
	"restaurant_management/internal/models/converters"
	"restaurant_management/internal/models/dtos"
	"restaurant_management/internal/models/entities"
	"restaurant_management/internal/repositories"
)

type CategoryService interface {
	FindAll(request *dto.CategorySearch) ([]dto.CategoryResponse, *dto.Pagination, error)
	FindById(categoryId string) (*dto.CategoryResponse, error)
	Create(request *dto.CategoryCreateRequest) (*dto.CategoryResponse, error)
	Edit(request *dto.CategoryUpdateRequest) (*dto.CategoryResponse, error)
	Delete(categoryId string) error
}

type CategoryServiceImpl struct {
	categoryRepository repositories.CategoryRepository
	validator          *validator.Validate
}

func NewCategoryServiceImpl(categoryRepository repositories.CategoryRepository, validator *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: categoryRepository,
		validator:          validator,
	}
}

func (s CategoryServiceImpl) FindAll(request *dto.CategorySearch) ([]dto.CategoryResponse, *dto.Pagination, error) {
	var allCategory []dto.CategoryResponse

	categories, pagination, err := s.categoryRepository.FindAll(request.Filters, request.Page, request.Limit)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		allCategory = append(allCategory, *converters.ContactToResponse(&category))
	}

	return allCategory, pagination, nil
}

func (s CategoryServiceImpl) FindById(categoryId string) (*dto.CategoryResponse, error) {
	category, err := s.categoryRepository.FindById(categoryId)
	if err != nil {
		panic(err.Error())
	}

	converter := converters.ContactToResponse(category)
	return converter, nil
}

func (s CategoryServiceImpl) Create(request *dto.CategoryCreateRequest) (*dto.CategoryResponse, error) {
	if err := s.validator.Struct(&request); err != nil {
		panic("request error")
	}

	category := &entities.Category{
		Name: request.Name,
	}

	newCategory, err := s.categoryRepository.Save(category)
	if err != nil {
		panic(err.Error())
	}

	converter := converters.ContactToResponse(newCategory)
	return converter, nil
}

func (s CategoryServiceImpl) Edit(request *dto.CategoryUpdateRequest) (*dto.CategoryResponse, error) {
	category, err := s.categoryRepository.FindById(request.ID)
	if err != nil {
		panic(err.Error())
	}

	if err := s.validator.Struct(&request); err != nil {
		panic("request error")
	}

	// updated data
	category.Name = request.Name

	updatedCategory, err := s.categoryRepository.Update(category)
	if err != nil {
		panic(err.Error())
	}

	converter := converters.ContactToResponse(updatedCategory)
	return converter, nil
}

func (s CategoryServiceImpl) Delete(categoryId string) error {
	category, err := s.categoryRepository.FindById(categoryId)
	if err != nil {
		panic(err.Error())
	}

	if err := s.categoryRepository.Delete(category.ID); err != nil {
		panic(err.Error())
	}

	return nil
}
