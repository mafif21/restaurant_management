package controllers

import (
	"net/http"
	"restaurant_management/internal/config"
	dto "restaurant_management/internal/models/dtos"
	"restaurant_management/internal/services"
	"strconv"
)

type CategoryController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Remove(w http.ResponseWriter, r *http.Request)
}

type CategoryControllerImpl struct {
	categoryService services.CategoryService
}

func NewCategoryControllerImpl(categoryservice services.CategoryService) CategoryController {
	return &CategoryControllerImpl{categoryService: categoryservice}
}

func (c CategoryControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	request := dto.CategorySearch{
		Filters: map[string]any{
			"name": name,
		},
		Page:  page,
		Limit: limit,
	}

	categories, pagination, err := c.categoryService.FindAll(&request)
	if err != nil {
		panic(err)
	}

	response := dto.Response{
		Status:  http.StatusOK,
		Message: "success get all categories",
		Data: map[string]any{
			"categories": categories,
			"paging":     pagination,
		},
	}

	config.WriteToBodyResponse(w, &response)
}

func (c CategoryControllerImpl) GetById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryControllerImpl) Remove(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
