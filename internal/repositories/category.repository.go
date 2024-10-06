package repositories

import (
	"errors"
	"gorm.io/gorm"
	"restaurant_management/internal/models/dtos"
	"restaurant_management/internal/models/entities"
)

type CategoryRepository interface {
	FindAll(filters map[string]any, page int, limit int) ([]entities.Category, *dto.Pagination, error)
	FindById(categoryId string) (*entities.Category, error)
	Save(category *entities.Category) (*entities.Category, error)
	Update(category *entities.Category) (*entities.Category, error)
	Delete(categoryId string) error
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryImpl(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (r CategoryRepositoryImpl) FindAll(filters map[string]any, page int, limit int) ([]entities.Category, *dto.Pagination, error) {
	var categories []entities.Category

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	skip := (page - 1) * limit // offset data

	query := r.db.Model(entities.Category{})
	for key, value := range filters {
		switch key {
		case "name":
			query = query.Where("name LIKE ?", "%"+value.(string)+"%")
		}
	}

	if err := query.Limit(limit).Offset(skip).Find(&categories).Error; err != nil {
		return nil, nil, err
	}

	var totalItems int64
	if err := r.db.Model(&entities.Category{}).Count(&totalItems).Error; err != nil {
		return nil, nil, err
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	return categories, &dto.Pagination{
		Page:      page,
		TotalItem: int(totalItems),
		TotalPage: totalPages,
	}, nil

}

func (r CategoryRepositoryImpl) FindById(categoryId string) (*entities.Category, error) {
	var category *entities.Category

	if err := r.db.First(&category, "id = ?", categoryId).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return category, nil
}

func (r CategoryRepositoryImpl) Save(category *entities.Category) (*entities.Category, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return nil, errors.New("failed to save data")
	}

	return category, nil
}

func (r CategoryRepositoryImpl) Update(category *entities.Category) (*entities.Category, error) {
	if err := r.db.Updates(&category).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return category, nil
}

func (r CategoryRepositoryImpl) Delete(categoryId string) error {
	if err := r.db.Delete(entities.Category{}, "id = ?", categoryId).Error; err != nil {
		return errors.New("failed to delete data")
	}

	return nil
}
