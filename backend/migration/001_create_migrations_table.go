package migration

import "gorm.io/gorm"

type CreateMigrationsTable struct{}

func (m *CreateMigrationsTable) Name() string {
	return "CreateMigrationsTable"
}

func (m *CreateMigrationsTable) Up(db *gorm.DB) error {
	return db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL UNIQUE,
			batch INT NOT NULL,
			executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`).Error
}

func (m *CreateMigrationsTable) Down(db *gorm.DB) error {
	return db.Exec("DROP TABLE IF EXISTS migrations;").Error
}
