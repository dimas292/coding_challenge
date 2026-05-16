package migration

import "gorm.io/gorm"

type CreateCategoriesTable struct{}

func (m *CreateCategoriesTable) Name() string {
	return "CreateCategoriesTable"
}

func (m *CreateCategoriesTable) Up(db *gorm.DB) error {
	return db.Exec(`
		CREATE TABLE IF NOT EXISTS t_categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			color VARCHAR(7) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`).Error
}

func (m *CreateCategoriesTable) Down(db *gorm.DB) error {
	return db.Exec("DROP TABLE IF EXISTS t_categories;").Error
}
