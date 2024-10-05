package dto

type Pagination struct {
	Page      int `json:"page"`
	TotalItem int `json:"total_item"`
	TotalPage int `json:"total_page"`
}
