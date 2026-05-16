package migration

import "gorm.io/gorm"

type CreateTodosTable struct{}

func (m *CreateTodosTable) Name() string {
	return "CreateTodosTable"
}

func (m *CreateTodosTable) Up(db *gorm.DB) error {
	return db.Exec(`
		CREATE TABLE IF NOT EXISTS t_todos (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			category_id INT REFERENCES t_categories(id) ON DELETE SET NULL,
			priority VARCHAR(50) DEFAULT 'MEDIUM',
			completed BOOLEAN DEFAULT false,
			due_date TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`).Error
}

func (m *CreateTodosTable) Down(db *gorm.DB) error {
	return db.Exec("DROP TABLE IF EXISTS t_todos;").Error
}
