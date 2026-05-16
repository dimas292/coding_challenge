package database

import (
	"backend-coding-challenge/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {
	var count int64
	db.Model(&model.Category{}).Count(&count)
	if count > 0 {
		fmt.Println("categories already seeded, skipping...")
		return
	}

	createdAt := time.Date(2024, 8, 1, 0, 0, 0, 0, time.FixedZone("WIB", 7*3600))

	categories := []model.Category{
		{Name: "Work", Color: "#3B82F6", CreatedAt: createdAt},
		{Name: "Hobbies", Color: "#1E104E", CreatedAt: createdAt},
		{Name: "Shopping", Color: "#B500B2", CreatedAt: createdAt},
		{Name: "Study", Color: "#6367FF", CreatedAt: createdAt},
		{Name: "Gaming", Color: "#008BFF", CreatedAt: createdAt},
	}

	if err := db.Create(&categories).Error; err != nil {
		fmt.Printf("failed to seed categories: %v\n", err)
		return
	}

	fmt.Printf("seeded %d categories successfully\n", len(categories))
}

func SeedTodos(db *gorm.DB) {
	var count int64
	db.Model(&model.Todo{}).Count(&count)
	if count > 0 {
		fmt.Println("todos already seeded, skipping...")
		return
	}

	parseTime := func(value string) time.Time {
		parsed, err := time.Parse(time.RFC3339, value)
		if err != nil {
			fmt.Printf("failed to parse time %q: %v\n", value, err)
			return time.Time{}
		}
		return parsed
	}

	var categories []model.Category
	if err := db.Find(&categories).Error; err != nil {
		fmt.Printf("failed to load categories: %v\n", err)
		return
	}

	categoryByName := make(map[string]int, len(categories))
	for _, category := range categories {
		categoryByName[category.Name] = category.ID
	}

	lookupCategoryID := func(name string) (int, bool) {
		id, ok := categoryByName[name]
		return id, ok
	}

	buildTodo := func(title, description, categoryName, priority, dueDate, createdAt string, completed bool) (model.Todo, bool) {
		categoryID, ok := lookupCategoryID(categoryName)
		if !ok {
			fmt.Printf("category %q not found for todo %q\n", categoryName, title)
			return model.Todo{}, false
		}

		createdAtTime := parseTime(createdAt)
		return model.Todo{
			Title:       title,
			Description: description,
			Completed:   completed,
			CategoryID:  categoryID,
			Priority:    priority,
			DueDate:     parseTime(dueDate),
			CreatedAt:   &createdAtTime,
		}, true
	}

	todos := []model.Todo{}
	seededTodos := []struct {
		title        string
		description  string
		categoryName string
		priority     string
		dueDate      string
		createdAt    string
		completed    bool
	}{
		{"Buy a new bag", "Go to shop for buying the bag", "Shopping", "low", "2024-08-03T23:59:59Z", "2024-07-31T10:00:00Z", false},
		{"Complete project proposal", "Draft and submit the Q3 project proposal to the manager", "Work", "high", "2024-08-05T17:00:00Z", "2024-07-31T10:30:00Z", true},
		{"Grocery shopping", "Buy vegetables, milk, and eggs for the week", "Shopping", "medium", "2024-08-02T18:00:00Z", "2024-07-31T11:00:00Z", false},
		{"Schedule dentist appointment", "Call the clinic to schedule routine checkup", "Hobbies", "low", "2024-08-10T12:00:00Z", "2024-07-31T14:20:00Z", false},
		{"Fix login bug", "Investigate and resolve the authentication issue on mobile app", "Work", "high", "2024-08-01T23:59:59Z", "2024-07-31T15:00:00Z", false},
		{"Pay internet bill", "Pay the monthly fiber optic subscription via mobile banking", "Hobbies", "high", "2024-08-05T23:59:59Z", "2024-08-01T08:00:00Z", true},
		{"Read new book", "Finish the first 3 chapters of the new sci-fi novel", "Hobbies", "low", "2024-08-07T22:00:00Z", "2024-08-01T09:00:00Z", false},
		{"Car maintenance", "Take the car to the workshop for oil change", "Hobbies", "medium", "2024-08-04T15:00:00Z", "2024-08-01T10:15:00Z", false},
		{"Update resume", "Add the latest project experiences and skills to the CV", "Hobbies", "medium", "2024-08-15T23:59:59Z", "2024-08-02T08:30:00Z", false},
		{"Clean the apartment", "Vacuum the living room and clean the kitchen", "Hobbies", "low", "2024-08-02T12:00:00Z", "2024-08-02T09:00:00Z", true},
		{"Review pull requests", "Review and merge pending pull requests from the backend team", "Work", "high", "2024-08-03T17:00:00Z", "2024-08-02T10:00:00Z", false},
		{"Call parents", "Catch up with family over the weekend", "Hobbies", "medium", "2024-08-04T20:00:00Z", "2024-08-02T11:00:00Z", false},
		{"Plan weekend trip", "Book hotels and look for attractions for the short trip", "Hobbies", "low", "2024-08-08T23:59:59Z", "2024-08-03T08:00:00Z", false},
		{"Renew domain name", "Pay for the personal portfolio website domain renewal", "Hobbies", "high", "2024-08-06T23:59:59Z", "2024-08-03T09:30:00Z", false},
		{"Water the plants", "Water all indoor and balcony plants", "Hobbies", "medium", "2024-08-03T10:00:00Z", "2024-08-03T07:00:00Z", true},
		{"Prepare presentation slides", "Create slides for the upcoming monthly review meeting", "Work", "high", "2024-08-05T09:00:00Z", "2024-08-03T11:00:00Z", false},
		{"Buy a birthday gift", "Find a suitable present for Alex's birthday party", "Shopping", "medium", "2024-08-07T18:00:00Z", "2024-08-03T13:00:00Z", false},
		{"Backup server data", "Run the manual backup script and verify the snapshot", "Work", "high", "2024-08-04T02:00:00Z", "2024-08-03T15:00:00Z", false},
		{"Do laundry", "Wash and fold clothes for the upcoming week", "Hobbies", "low", "2024-08-04T12:00:00Z", "2024-08-04T08:00:00Z", false},
		{"Watch tutorial video", "Watch the new course module on advanced CSS animations", "Work", "medium", "2024-08-06T20:00:00Z", "2024-08-04T09:30:00Z", false},
	}

	for _, item := range seededTodos {
		todo, ok := buildTodo(item.title, item.description, item.categoryName, item.priority, item.dueDate, item.createdAt, item.completed)
		if !ok {
			return
		}
		todos = append(todos, todo)
	}

	if err := db.Create(&todos).Error; err != nil {
		fmt.Printf("failed to seed todos: %v\n", err)
		return
	}

	fmt.Printf("seeded %d todos successfully\n", len(todos))
}
