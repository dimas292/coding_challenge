package migration

import "gorm.io/gorm"

type AddIndexes struct{}

func (m *AddIndexes) Name() string {
	return "AddIndexes"
}

func (m *AddIndexes) Up(db *gorm.DB) error {
	// Create indexes for better query performance
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_todos_category_id ON t_todos(category_id);",
		"CREATE INDEX IF NOT EXISTS idx_todos_completed ON t_todos(completed);",
		"CREATE INDEX IF NOT EXISTS idx_todos_created_at ON t_todos(created_at);",
		"CREATE INDEX IF NOT EXISTS idx_todos_priority ON t_todos(priority);",
		"CREATE INDEX IF NOT EXISTS idx_categories_name ON t_categories(name);",
	}

	for _, idx := range indexes {
		if err := db.Exec(idx).Error; err != nil {
			return err
		}
	}

	return nil
}

func (m *AddIndexes) Down(db *gorm.DB) error {
	// Drop indexes
	indexes := []string{
		"DROP INDEX IF EXISTS idx_todos_category_id;",
		"DROP INDEX IF EXISTS idx_todos_completed;",
		"DROP INDEX IF EXISTS idx_todos_created_at;",
		"DROP INDEX IF EXISTS idx_todos_priority;",
		"DROP INDEX IF EXISTS idx_categories_name;",
	}

	for _, idx := range indexes {
		if err := db.Exec(idx).Error; err != nil {
			return err
		}
	}

	return nil
}
