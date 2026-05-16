package model

import "time"

type Category struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Color     string    `gorm:"column:color" json:"color"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Category) TableName() string {
	return "t_categories"
}