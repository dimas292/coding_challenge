# Coding Challenge Todo App

## Table of Contents

1. [Introduction](#introduction)
2. [Setup and Installation](#setup-and-installation)
3. [Architecture](#architecture)
   - [Folder Structure](#folder-structure)
   - [Components](#components)
   - [Program Flow](#program-flow)
4. [Frontend Documentation](#frontend-documentation)
   - [Application Structure](#application-structure)
   - [React Components](#react-components)
   - [Context API](#context-api)
   - [Custom Hooks](#custom-hooks)
   - [Services](#services)
   - [Utils and Constants](#utils-and-constants)
5. [API Documentation](#api-documentation)
   - [RESTful API](#restful-api)
   - [API Usage Examples](#api-usage-examples)
6. [Database Schema](#database-schema)
   - [Database Tables](#database-tables)
   - [Relations](#relations)
   - [Models](#models)
   - [Entity-Relationship Diagram](#entity-relationship-diagram)
7. [Database Migrations](#database-migrations)
   - [Creating a New Migration](#creating-a-new-migration)
   - [Migration Examples](#migration-examples)

---

## Introduction

Coding Challenge is a modern Todo Management System application. It enables users to create, manage, and track task lists with categories and priorities. The system consists of a backend API using Go and Gin Framework, a frontend built with React and TypeScript using Ant Design UI, and a PostgreSQL database.

---

## Setup and Installation

### Prerequisites

- Docker and Docker Compose
- Node.js v20+ (for local development)
- Go v1.26+ (for local development)
- PostgreSQL 16 (for local development)

### Setup with Docker Compose

```bash
# Clone repository
git clone https://github.com/dimas292/coding_challenge.git
cd coding_challenge

# Build and run containers
docker compose build
docker compose up -d

# Access the application
# Frontend: http://localhost:5173
# Backend: http://localhost:4444
# Database: localhost:5432
```

### Local Development Setup

```bash
# Backend
cd backend
go mod download
go run main.go

# Frontend (new terminal)
cd frontend
npm install
npm run dev
```

### Environment Variables (.env)

Store sensitive data in `.env` file (don't commit). Example `.env` file:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=todo_db
VITE_API_URL=http://localhost:4444/api
```

Use [.env.example](.env.example) as a template, then copy to `.env`.

---

## Architecture

### Folder Structure

```
coding_challenge/
├── backend/
│   ├── config/          # Application configuration
│   ├── database/        # Database connection and seeding
│   ├── handler/         # Request handlers
│   ├── model/           # Data models
│   ├── repository/      # Data access layer
│   ├── router/          # Route definitions
│   ├── server/          # Server setup
│   ├── service/         # Business logic
│   ├── utils/           # Utility functions
│   ├── sql/             # Migration files
│   ├── main.go          # Entry point
│   ├── go.mod/go.sum    # Dependency management
│   └── Dockerfile       # Container image
├── frontend/
│   ├── src/
│   │   ├── components/  # React components
│   │   ├── context/     # Context API
│   │   ├── hooks/       # Custom hooks
│   │   ├── services/    # API services
│   │   ├── utils/       # Utility functions
│   │   └── main.tsx     # Entry point
│   ├── package.json
│   └── Dockerfile
├── docker-compose.yaml  # Docker Compose configuration
└── README.md
```

### Components

**Backend:**

- **Gin Framework**: REST API web framework
- **GORM**: ORM for database management
- **PostgreSQL**: Relational database
- **CORS**: Cross-Origin Resource Sharing

**Frontend:**

- **React 19**: UI library
- **TypeScript**: Type-safe JavaScript
- **Vite**: Build tool
- **Ant Design**: UI component library
- **React Icons**: Icon library

**Database:**

- **PostgreSQL 16**: Database server
- **Docker Volume**: Persistent data storage

### Program Flow

```
1. User opens Frontend (React)
2. Frontend sends request to Backend API (http://localhost:4444/api)
3. Backend (Gin) receives request and validates
4. Handler processes business logic via Service
5. Service accesses data via Repository
6. Repository queries PostgreSQL
7. Data returned to Frontend as JSON response
8. Frontend updates UI based on response
```

---

## Frontend Documentation

### Application Structure

```
src/
├── App.tsx              # Main application component
├── main.tsx             # Application entry point
├── index.css            # Global styles
├── App.css              # App component styles
├── components/
│   ├── ui/              # UI components
│   │   ├── TodoHeader.tsx       # Header with add button
│   │   ├── SearchBar.tsx        # Search input
│   │   └── FilterBar.tsx        # Filter by category/priority/status
│   ├── DataTable.tsx    # Main table displaying todos
│   ├── TableColumns.tsx # Column definitions for table
│   ├── CreateTodo.tsx   # Create todo component
│   ├── EditTodo.tsx     # Edit todo component
│   └── DrawerForm.tsx   # Form drawer for create/edit
├── context/
│   └── TodoContext.tsx  # Context for state management
├── hooks/
│   ├── useTodoFetch.ts  # Hook for fetching todos
│   └── useFormSubmit.ts # Hook for form submission
├── services/
│   ├── TodoService.ts   # API calls for todos
│   └── CategoryService.ts # API calls for categories
└── utils/
    ├── constants.ts     # Application constants
    ├── dateUtils.ts     # Date manipulation utilities
    └── tagRenderer.tsx  # Utility for rendering tags/badges
```

### React Components

#### 1. **App.tsx** (Main Component)

Main component that manages layout and overall application state:

- Manages pagination state
- Manages filter state (search, priority, category, completed)
- Calls `useTodoFetch` hook to fetch data
- Renders header, search bar, filter bar, and data table

```tsx
// Using Context
const { openDrawer, todosVersion, todos } = UseTodo();
const { loading, fetchTodos } = useTodoFetch();
```

#### 2. **DataTable.tsx** (Table Component)

Displays todos in table format with features:

- Pagination
- Sorting
- Column definitions from `TableColumns.tsx`
- Row actions (edit/delete)

#### 3. **CreateTodo.tsx** (Create Dialog)

Wrapper component displaying drawer form in "create" mode:

- Fetches categories from API when drawer opens
- Passes category options to `DrawerForm`

#### 4. **EditTodo.tsx** (Edit Dialog)

Wrapper component displaying drawer form in "edit" mode:

- Fetches todo data to edit
- Populates form with existing data
- Submits updates to API

#### 5. **DrawerForm.tsx** (Form Component)

Reusable form component for create and edit:

- Inputs: title, description, category, priority, due_date
- Form validation
- Handles submit with `useFormSubmit` hook

#### 6. **UI Components** (`components/ui/`)

Additional UI components:

- **TodoHeader**: Header with "Add Todo" button
- **SearchBar**: Search input with debouncing
- **FilterBar**: Select for filtering priority, category, and status

### Context API

#### **TodoContext.tsx**

Global state management using Context API:

```tsx
interface Todo {
  id: number;
  title: string;
  description: string;
  category: Category;
  priority: string; // high, medium, low
  completed: boolean;
  due_date: string;
  created_at: string;
}

interface TodoContextType {
  todos: Todo[]; // List of todos
  setTodos: (todos: Todo[]) => void; // Set todos
  addTodo: (todo: Todo) => void; // Add todo to list
  deleteTodo: (id: number) => void; // Remove todo from list
  editTodo: (updatedTodo: Todo) => void; // Update todo in list
  refreshTodos: () => void; // Trigger re-fetch
  editingTodoId: number | null; // ID of todo being edited
  drawerOpen: boolean; // Drawer visibility
  openDrawer: (todoId?: number) => void; // Open drawer
  closeDrawer: () => void; // Close drawer
  categoryOption: { label: string; value: number }[];
  setCategoryOption: (opts: any[]) => void;
  priorityParam: (p: string) => string; // Helper functions
  categoryParam: (c: string) => string;
  completedParam: (s: string) => string;
}
```

Used with hook: `const context = UseTodo()`

### Custom Hooks

#### 1. **useTodoFetch.ts**

Hook for fetching todos from backend:

```tsx
const { loading, fetchTodos } = useTodoFetch();

// Usage:
const pagination = await fetchTodos(
  page, // Current page
  limit, // Items per page
  search, // Search keyword
  priority, // Filter priority
  category, // Filter category
  completed, // Filter status
);
// Returns: { current_page, per_page, total, total_pages }
```

**Features:**

- Prevents multiple simultaneous fetches using ref
- Error handling with notifications
- TypeScript generics for response typing

#### 2. **useFormSubmit.ts**

Hook for handling form submission:

```tsx
const { handleFormSubmit } = useFormSubmit();

// Usage:
await handleFormSubmit({
  form, // Ant Design Form instance
  mode, // "create" or "edit"
  initialData, // Data for edit mode
  onSuccess, // Callback after submit
});
```

**Features:**

- Calls CreateTodo or UpdateTodo based on mode
- Shows success/error notifications
- Calls onSuccess callback

### Services

#### 1. **TodoService.ts**

API calls for todo management:

```tsx
// Create new todo
CreateTodo(data: Todo): Promise<any>

// Update existing todo
UpdateTodo(id: number, data: Todo): Promise<any>

// Delete todo
DeleteTodo(id: number): Promise<any>
```

#### 2. **CategoryService.ts**

API calls for category management:

```tsx
// Get all categories
GetCategory(): Promise<{ data: Category[] }>

// Get category by ID
GetCategoryById(id: number): Promise<{ data: Category }>
```

**Base URL**: Uses environment variable `VITE_API_URL` (default: `http://localhost:4444/api`)

### Utils and Constants

#### **constants.ts**

Global application constants:

```tsx
// Priority options for filter and select
export const PRIORITY_OPTIONS = [
  { text: "Low", value: "low" },
  { text: "Medium", value: "medium" },
  { text: "High", value: "high" },
];

// Color mapping for priority
export const PRIORITY_COLOR_MAP = {
  high: "red",
  medium: "orange",
  low: "green",
};

// Status options
export const COMPLETED_OPTIONS = [
  { text: "Done", value: true },
  { text: "Pending", value: false },
];

// Form modes
export const FORM_MODE = {
  CREATE: "create",
  EDIT: "edit",
};
```

#### **dateUtils.ts**

Utility functions for date handling:

- Format date to display format
- Parse date from API response

#### **tagRenderer.tsx**

Utility for rendering tags and badges in table

---

## API Documentation

### RESTful API

#### Todo Endpoints

| Method | Endpoint         | Description                  |
| ------ | ---------------- | ---------------------------- |
| POST   | `/api/todos/`    | Create a new todo            |
| GET    | `/api/todos/`    | Get all todos (with filters) |
| GET    | `/api/todos/:id` | Get todo by ID               |
| PUT    | `/api/todos/:id` | Update todo                  |
| DELETE | `/api/todos/:id` | Delete todo                  |

#### Category Endpoints

| Method | Endpoint            | Description         |
| ------ | ------------------- | ------------------- |
| POST   | `/api/category/`    | Create new category |
| GET    | `/api/category/`    | Get all categories  |
| GET    | `/api/category/:id` | Get category by ID  |
| PUT    | `/api/category/`    | Update category     |
| DELETE | `/api/category/:id` | Delete category     |

#### Query Parameters

**GET /api/todos/**

```
?page=1                 # Page number (default: 1)
&limit=10              # Items per page (default: 10)
&search=keyword        # Search by title/description
&priority=high         # Filter by priority (high/medium/low)
&category=1            # Filter by category ID
&completed=true        # Filter by completion status
```

### API Usage Examples

#### 1. Create a New Todo

```bash
curl -X POST http://localhost:4444/api/todos/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Learn Gin Framework",
    "category_id": 1,
    "priority": "high",
    "due_date": "2026-05-20T00:00:00Z"
  }'
```

**Response:**

```json
{
  "message": "Todo created successfully",
  "data": {
    "id": 1,
    "title": "Learn Go",
    "description": "Learn Gin Framework",
    "category_id": 1,
    "priority": "high",
    "completed": false,
    "due_date": "2026-05-20T00:00:00Z",
    "created_at": "2026-05-16T11:30:00Z"
  }
}
```

#### 2. Get All Todos with Filters

```bash
curl -X GET "http://localhost:4444/api/todos/?page=1&limit=10&priority=high&completed=false"
```

#### 3. Update Todo

```bash
curl -X PUT http://localhost:4444/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Advanced Go",
    "completed": true
  }'
```

#### 4. Delete Todo

```bash
curl -X DELETE http://localhost:4444/api/todos/1
```

#### 5. Create Category

```bash
curl -X POST http://localhost:4444/api/category/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Learning",
    "color": "#FF5733"
  }'
```

---

## Database Schema

### Database Tables

#### t_todos

Main table storing todo data.

| Column      | Type      | Constraint                  | Description                |
| ----------- | --------- | --------------------------- | -------------------------- |
| id          | INT       | PRIMARY KEY, AUTO_INCREMENT | Unique identifier          |
| title       | VARCHAR   | NOT NULL                    | Todo title                 |
| description | TEXT      |                             | Detailed description       |
| category_id | INT       | FOREIGN KEY                 | Reference to t_categories  |
| priority    | VARCHAR   |                             | Priority (high/medium/low) |
| completed   | BOOLEAN   | DEFAULT false               | Completion status          |
| due_date    | TIMESTAMP |                             | Deadline date              |
| created_at  | TIMESTAMP | DEFAULT NOW()               | Creation timestamp         |

#### t_categories

Table storing category data.

| Column     | Type      | Constraint                  | Description        |
| ---------- | --------- | --------------------------- | ------------------ |
| id         | INT       | PRIMARY KEY, AUTO_INCREMENT | Unique identifier  |
| name       | VARCHAR   | NOT NULL                    | Category name      |
| color      | VARCHAR   |                             | Hex color code     |
| created_at | TIMESTAMP | DEFAULT NOW()               | Creation timestamp |

### Relations

**One-to-Many Relationship:**

- One Category has many Todos
- Foreign Key: `t_todos.category_id` → `t_categories.id`

### Models

**Todo Model (Go)**

```go
type Todo struct {
    ID          int       // Unique identifier
    Title       string    // Todo title
    Description string    // Detailed description
    CategoryID  int       // Category reference
    Category    Category  // Embedded category object
    Priority    string    // high, medium, low
    Completed   bool      // Completion status
    DueDate     time.Time // Deadline date
    CreatedAt   *time.Time // Creation timestamp
}
```

**Category Model (Go)**

```go
type Category struct {
    ID        int       // Unique identifier
    Name      string    // Category name
    Color     string    // Hex color code
    CreatedAt time.Time // Creation timestamp
}
```

**Todo Model (TypeScript/Frontend)**

```tsx
interface Todo {
  id: number;
  title: string;
  description: string;
  category: Category;
  priority: string; // high, medium, low
  completed: boolean;
  due_date: string; // ISO 8601 format
  created_at: string; // ISO 8601 format
}
```

### Entity-Relationship Diagram

![erd](https://res.cloudinary.com/dmx8hcmxh/image/upload/v1778906893/todo.drawio_lrvdht.png)

Relationship: One-to-Many

- One category can have many todos
- Each todo must have one category

---

## Database Migrations

### Overview

This project uses manual GORM migrations for database schema management. Each migration is tracked in the `migrations` table and runs only once.

**Migration Files Location:** `backend/migration/`

| File | Purpose |
|------|---------|
| `migration.go` | Core runner with idempotent logic and tracking |
| `001_create_migrations_table.go` | Creates `migrations` tracking table |
| `002_create_categories_table.go` | Creates `t_categories` table |
| `003_create_todos_table.go` | Creates `t_todos` table with foreign key |
| `004_add_indexes.go` | Adds performance indexes |

### How Migrations Work

1. **Idempotent Execution** - Each migration runs only once per deployment
2. **Tracked in Database** - Migration history stored in `migrations` table
3. **Reversible** - Every migration has `Up()` and `Down()` methods
4. **Ordered** - Executed sequentially in the order defined in `migration.go`

### Creating a New Migration

**Step 1:** Create a new file in `backend/migration/005_your_migration_name.go`

```go
package migration

import "gorm.io/gorm"

type YourMigrationName struct{}

func (m *YourMigrationName) Name() string {
    return "YourMigrationName"
}

func (m *YourMigrationName) Up(db *gorm.DB) error {
    return db.Exec(`
        ALTER TABLE t_todos ADD COLUMN IF NOT EXISTS tags VARCHAR(255);
    `).Error
}

func (m *YourMigrationName) Down(db *gorm.DB) error {
    return db.Exec(`
        ALTER TABLE t_todos DROP COLUMN IF EXISTS tags;
    `).Error
}
```

**Step 2:** Register in `backend/migration/migration.go`

```go
var migrations = []Migration{
    &CreateCategoriesTable{},
    &CreateTodosTable{},
    &CreateMigrationsTable{},
    &AddIndexes{},
    &YourMigrationName{},  // Add here
}
```

**Step 3:** Rebuild and run

```bash
cd backend
go build -o server
./server  # Migrations run automatically on startup
```

### Migration Examples

**Example 1: Add a new column with default value**

```go
func (m *AddStatusColumn) Up(db *gorm.DB) error {
    return db.Exec(`
        ALTER TABLE t_todos 
        ADD COLUMN IF NOT EXISTS status VARCHAR(50) DEFAULT 'pending';
    `).Error
}

func (m *AddStatusColumn) Down(db *gorm.DB) error {
    return db.Exec(`
        ALTER TABLE t_todos DROP COLUMN IF EXISTS status;
    `).Error
}
```

**Example 2: Create a new table**

```go
func (m *CreateCommentsTable) Up(db *gorm.DB) error {
    return db.Exec(`
        CREATE TABLE IF NOT EXISTS t_comments (
            id SERIAL PRIMARY KEY,
            todo_id INT NOT NULL REFERENCES t_todos(id) ON DELETE CASCADE,
            content TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
        
        CREATE INDEX IF NOT EXISTS idx_comments_todo_id ON t_comments(todo_id);
    `).Error
}

func (m *CreateCommentsTable) Down(db *gorm.DB) error {
    return db.Exec("DROP TABLE IF EXISTS t_comments;").Error
}
```

**Example 3: Add foreign key constraint**

```go
func (m *AddForeignKey) Up(db *gorm.DB) error {
    return db.Exec(`
        ALTER TABLE t_todos 
        ADD CONSTRAINT fk_todos_category 
        FOREIGN KEY (category_id) 
        REFERENCES t_categories(id) ON DELETE SET NULL;
    `).Error
}

func (m *AddForeignKey) Down(db *gorm.DB) error {
    return db.Exec(`
        ALTER TABLE t_todos 
        DROP CONSTRAINT IF EXISTS fk_todos_category;
    `).Error
}
```

### Best Practices

✅ **DO:**
- Use `CREATE TABLE IF NOT EXISTS` and `DROP TABLE IF EXISTS`
- Always implement both `Up()` and `Down()` methods
- Use descriptive migration names
- Test migrations locally before deploying
- Keep migrations focused on a single change

❌ **DON'T:**
- Skip the `Down()` method
- Hardcode IDs or specific values
- Make migrations too large with multiple unrelated changes
- Forget to update the `migrations` array in `migration.go`

### Checking Migration Status

Query the migrations table to see execution history:

```sql
SELECT * FROM migrations ORDER BY id DESC;
```

Output:
```
 id |         name          | batch |        executed_at
----+-----------------------+-------+----------------------------
  3 | AddIndexes            |     3 | 2026-05-16 06:42:31
  2 | CreateTodosTable      |     2 | 2026-05-16 06:42:31
  1 | CreateCategoriesTable |     1 | 2026-05-16 06:42:31
```

---

```
