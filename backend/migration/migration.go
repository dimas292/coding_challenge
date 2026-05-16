package migration

import (
	"fmt"
	"gorm.io/gorm"
)

type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
	Name() string
}

var migrations = []Migration{
	&CreateCategoriesTable{},
	&CreateTodosTable{},
	&CreateMigrationsTable{},
	&AddIndexes{},
}

func RunAll(db *gorm.DB) error {
	fmt.Println("[Migration] Creating migrations table...")
	if err := createMigrationsTable(db); err != nil {
		return err
	}

	for _, m := range migrations {
		// Skip migrations table itself
		if m.Name() == "CreateMigrationsTable" {
			continue
		}

		// Check if migration already ran
		if migrationExists(db, m.Name()) {
			fmt.Printf("[Migration] Skipping %s (already ran)\n", m.Name())
			continue
		}

		fmt.Printf("[Migration] Running %s...\n", m.Name())
		if err := m.Up(db); err != nil {
			fmt.Printf("[Migration] ❌ Error running %s: %v\n", m.Name(), err)
			return err
		}

		// Record migration
		if err := recordMigration(db, m.Name()); err != nil {
			fmt.Printf("[Migration] Warning: Could not record %s: %v\n", m.Name(), err)
		}

		fmt.Printf("[Migration] ✓ %s completed\n", m.Name())
	}

	fmt.Println("[Migration] All migrations completed successfully!")
	return nil
}

func RollbackAll(db *gorm.DB) error {
	fmt.Println("[Migration] Rolling back all migrations...")
	for i := len(migrations) - 1; i >= 0; i-- {
		m := migrations[i]
		if m.Name() == "CreateMigrationsTable" {
			continue
		}

		fmt.Printf("[Migration] Rolling back %s...\n", m.Name())
		if err := m.Down(db); err != nil {
			fmt.Printf("[Migration] ❌ Error rolling back %s: %v\n", m.Name(), err)
			return err
		}

		// Remove migration record
		if err := removeMigration(db, m.Name()); err != nil {
			fmt.Printf("[Migration] Warning: Could not remove %s record: %v\n", m.Name(), err)
		}

		fmt.Printf("[Migration] ✓ %s rolled back\n", m.Name())
	}

	fmt.Println("[Migration] All rollbacks completed!")
	return nil
}

func createMigrationsTable(db *gorm.DB) error {
	return db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL UNIQUE,
			batch INT NOT NULL,
			executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`).Error
}

func migrationExists(db *gorm.DB, name string) bool {
	var count int64
	db.Raw("SELECT COUNT(*) FROM migrations WHERE name = ?", name).Scan(&count)
	return count > 0
}

func recordMigration(db *gorm.DB, name string) error {
	var batch int64 = 1
	db.Raw("SELECT COALESCE(MAX(batch), 0) + 1 FROM migrations").Scan(&batch)
	return db.Exec("INSERT INTO migrations (name, batch) VALUES (?, ?)", name, batch).Error
}

func removeMigration(db *gorm.DB, name string) error {
	return db.Exec("DELETE FROM migrations WHERE name = ?", name).Error
}

func GetStatus(db *gorm.DB) error {
	fmt.Println("\n[Migration] Status:")
	fmt.Println("==================")

	var results []map[string]interface{}
	db.Raw("SELECT name, executed_at FROM migrations ORDER BY id DESC").Scan(&results)

	if len(results) == 0 {
		fmt.Println("No migrations have been run yet.")
		return nil
	}

	for _, row := range results {
		fmt.Printf("✓ %s (%s)\n", row["name"], row["executed_at"])
	}

	fmt.Println("==================\n")
	return nil
}
