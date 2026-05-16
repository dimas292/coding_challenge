package utils

type FormatTodoResponse struct {
	Data interface{} `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage int `json:"per_page"`
	Total int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
