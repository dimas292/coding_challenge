package model

import "time"

type Todo struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CategoryID  int      `gorm:"column:category_id" json:"category_id"`
	Category    Category `json:"category"`
	Priority    string   `gorm:"column:priority" json:"priority"`
	Completed   bool     `gorm:"column:completed" json:"completed"`
	DueDate     time.Time `gorm:"column:due_date" json:"due_date"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Todo) TableName() string {
	return "t_todos"
}

