package utils

type TodoFilterParam struct {
	Page int64
	Limit int64
	Category int64
	Completed *bool
	Priority string
	SortBy string 
	OrderBy string
	Search string
}