# Coding Challenge Todo App

## Daftar Isi

1. [Pendahuluan](#pendahuluan)
2. [Persiapan dan Instalasi](#persiapan-dan-instalasi)
3. [Arsitektur](#arsitektur)
   - [Struktur Folder](#struktur-folder)
   - [Komponen](#komponen)
   - [Alur Program](#alur-program)
4. [Dokumentasi Frontend](#dokumentasi-frontend)
   - [Struktur Aplikasi](#struktur-aplikasi)
   - [Komponen React](#komponen-react)
   - [Context API](#context-api)
   - [Custom Hooks](#custom-hooks)
   - [Services](#services)
   - [Utils dan Constants](#utils-dan-constants)
5. [Dokumentasi API](#dokumentasi-api)
   - [RESTful API](#restful-api)
   - [Contoh Penggunaan API](#contoh-penggunaan-api)
6. [Database Skema](#database-skema)
   - [Tabel Database](#tabel-database)
   - [Relasi](#relasi)
   - [Model](#model)
   - [Entity-Relationship Diagram](#entity-relationship-diagram)

---

## Pendahuluan

Coding Challenge adalah aplikasi manajemen tugas (Todo Management System) yang dibangun dengan teknologi modern. Aplikasi ini memungkinkan pengguna untuk membuat, mengelola, dan melacak daftar tugas dengan kategori dan prioritas. Sistem ini terdiri dari backend API menggunakan Go dan Gin Framework, frontend menggunakan React UI Ant Design dengan TypeScript, serta database PostgreSQL.

---

## Persiapan dan Instalasi

### Prasyarat

- Docker dan Docker Compose
- Node.js v20+ (untuk development lokal)
- Go v1.26+ (untuk development lokal)
- PostgreSQL 16 (untuk development lokal)

### Setup dengan Docker Compose

```bash
# Clone repository
git clone https://github.com/dimas292/coding_challenge.git
cd coding_challenge

# Build dan jalankan container
docker compose build
docker compose up -d

# Akses aplikasi
# Frontend: http://localhost:5173
# Backend: http://localhost:4444
# Database: localhost:5432
```

### Setup Lokal Development

```bash
# Backend
cd backend
go mod download
go run main.go

# Frontend (terminal baru)
cd frontend
npm install
npm run dev
```

---

## Arsitektur

### Struktur Folder

```
coding_challenge/
├── backend/
│   ├── config/          # Konfigurasi aplikasi
│   ├── database/        # Koneksi dan seed database
│   ├── handler/         # Request handler
│   ├── model/           # Data model
│   ├── repository/      # Data access layer
│   ├── router/          # Route definition
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
├── docker-compose.yaml  # Compose configuration
└── README.md
```

### Komponen

**Backend:**

- **Gin Framework**: REST API web framework
- **GORM**: ORM untuk database management
- **PostgreSQL**: Database relasional
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

### Alur Program

```
1. User membuka Frontend (React)
2. Frontend mengirim request ke Backend API (http://localhost:4444/api)
3. Backend (Gin) menerima request dan validasi
4. Handler memproses logika bisnis via Service
5. Service mengakses data via Repository
6. Repository query ke PostgreSQL
7. Data dikembalikan ke Frontend dengan JSON response
8. Frontend update UI berdasarkan response
```

---

## Frontend

### Struktur Aplikasi

```
src/
├── App.tsx              # Main application component
├── main.tsx             # Application entry point
├── index.css            # Global styles
├── App.css              # App component styles
├── components/
│   ├── ui/              # UI components
│   │   ├── TodoHeader.tsx       # Header dengan tombol tambah
│   │   ├── SearchBar.tsx        # Search input
│   │   └── FilterBar.tsx        # Filter by category/priority/status
│   ├── DataTable.tsx    # Main table untuk menampilkan todos
│   ├── TableColumns.tsx # Column definition untuk table
│   ├── CreateTodo.tsx   # Komponen untuk create todo
│   ├── EditTodo.tsx     # Komponen untuk edit todo
│   └── DrawerForm.tsx   # Form drawer untuk create/edit
├── context/
│   └── TodoContext.tsx  # Context untuk state management
├── hooks/
│   ├── useTodoFetch.ts  # Hook untuk fetch todos dari API
│   └── useFormSubmit.ts # Hook untuk submit form
├── services/
│   ├── TodoService.ts   # API calls untuk todo
│   └── CategoryService.ts # API calls untuk category
└── utils/
    ├── constants.ts     # Konstanta aplikasi
    ├── dateUtils.ts     # Date manipulation utilities
    └── tagRenderer.tsx  # Utility untuk render tag/badge
```

### Komponen React

#### 1. **App.tsx** (Main Component)

Komponen utama yang mengatur layout dan state keseluruhan aplikasi:

- Mengelola pagination state
- Mengelola filter state (search, priority, category, completed)
- Memanggil `useTodoFetch` hook untuk fetch data
- Render header, search bar, filter bar, dan data table

```tsx
// Menggunakan Context
const { openDrawer, todosVersion, todos } = UseTodo();
const { loading, fetchTodos } = useTodoFetch();
```

#### 2. **DataTable.tsx** (Table Component)

Menampilkan daftar todos dalam format table dengan fitur:

- Pagination
- Sorting
- Column definition dari `TableColumns.tsx`
- Row actions (edit/delete)

#### 3. **CreateTodo.tsx** (Create Dialog)

Wrapper component untuk menampilkan drawer form ketika mode adalah "create":

- Fetch kategori dari API saat drawer dibuka
- Pass kategori options ke `DrawerForm`

#### 4. **EditTodo.tsx** (Edit Dialog)

Wrapper component untuk menampilkan drawer form ketika mode adalah "edit":

- Fetch data todo yang akan diedit
- Populate form dengan data existing
- Submit update ke API

#### 5. **DrawerForm.tsx** (Form Component)

Komponen form reusable untuk create dan edit:

- Input: title, description, category, priority, due_date
- Validasi form
- Handle submit dengan `useFormSubmit` hook

#### 6. **UI Components** (`components/ui/`)

Komponen UI tambahan:

- **TodoHeader**: Header dengan tombol "Add Todo"
- **SearchBar**: Input pencarian dengan debouncing
- **FilterBar**: Select untuk filter priority, category, dan status

### Context API

#### **TodoContext.tsx**

State management global menggunakan Context API:

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
  todos: Todo[]; // List todos
  setTodos: (todos: Todo[]) => void; // Set todos
  addTodo: (todo: Todo) => void; // Add todo ke list
  deleteTodo: (id: number) => void; // Remove todo dari list
  editTodo: (updatedTodo: Todo) => void; // Update todo di list
  refreshTodos: () => void; // Trigger re-fetch
  editingTodoId: number | null; // ID todo yang sedang di-edit
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

Digunakan dengan hook: `const context = UseTodo()`

### Custom Hooks

#### 1. **useTodoFetch.ts**

Hook untuk fetch data todos dari backend:

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

**Fitur:**

- Prevent multiple simultaneous fetches dengan ref
- Error handling dengan notification
- TypeScript generics untuk response typing

#### 2. **useFormSubmit.ts**

Hook untuk handle form submission:

```tsx
const { handleFormSubmit } = useFormSubmit();

// Usage:
await handleFormSubmit({
  form, // Ant Design Form instance
  mode, // "create" atau "edit"
  initialData, // Data untuk edit mode
  onSuccess, // Callback setelah submit
});
```

**Fitur:**

- Call CreateTodo atau UpdateTodo sesuai mode
- Show success/error notification
- Call onSuccess callback

### Services

#### 1. **TodoService.ts**

API calls untuk todo management:

```tsx
// Create new todo
CreateTodo(data: Todo): Promise<any>

// Update existing todo
UpdateTodo(id: number, data: Todo): Promise<any>

// Delete todo
DeleteTodo(id: number): Promise<any>
```

#### 2. **CategoryService.ts**

API calls untuk category management:

```tsx
// Get all categories
GetCategory(): Promise<{ data: Category[] }>

// Get category by ID
GetCategoryById(id: number): Promise<{ data: Category }>
```

**Base URL**: Menggunakan environment variable `VITE_API_URL` (default: `http://localhost:4444/api`)

### Utils dan Constants

#### **constants.ts**

Konstanta global aplikasi:

```tsx
// Priority options untuk filter dan select
export const PRIORITY_OPTIONS = [
  { text: "Low", value: "low" },
  { text: "Medium", value: "medium" },
  { text: "High", value: "high" },
];

// Color mapping untuk priority
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

Utility functions untuk date handling:

- Format date ke display format
- Parse date dari API response

#### **tagRenderer.tsx**

Utility untuk render tag dan badge di table

---

## Dokumentasi API

### RESTful API

#### Endpoints Todo

| Method | Endpoint         | Deskripsi                           |
| ------ | ---------------- | ----------------------------------- |
| POST   | `/api/todos/`    | Buat todo baru                      |
| GET    | `/api/todos/`    | Dapatkan semua todo (dengan filter) |
| GET    | `/api/todos/:id` | Dapatkan todo by ID                 |
| PUT    | `/api/todos/:id` | Update todo                         |
| DELETE | `/api/todos/:id` | Hapus todo                          |

#### Endpoints Category

| Method | Endpoint            | Deskripsi               |
| ------ | ------------------- | ----------------------- |
| POST   | `/api/category/`    | Buat kategori baru      |
| GET    | `/api/category/`    | Dapatkan semua kategori |
| GET    | `/api/category/:id` | Dapatkan kategori by ID |
| PUT    | `/api/category/`    | Update kategori         |
| DELETE | `/api/category/:id` | Hapus kategori          |

#### Query Parameters

**GET /api/todos/**

```
?page=1                 # Halaman (default: 1)
&limit=10              # Items per page (default: 10)
&search=keyword        # Cari berdasarkan title/description
&priority=high         # Filter by priority (high/medium/low)
&category=1            # Filter by category ID
&completed=true        # Filter by completion status
```

### Contoh Penggunaan API

#### 1. Buat Todo Baru

```bash
curl -X POST http://localhost:4444/api/todos/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Go",
    "description": "Belajar Gin Framework",
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
    "title": "Belajar Go",
    "description": "Belajar Gin Framework",
    "category_id": 1,
    "priority": "high",
    "completed": false,
    "due_date": "2026-05-20T00:00:00Z",
    "created_at": "2026-05-16T11:30:00Z"
  }
}
```

#### 2. Dapatkan Semua Todo dengan Filter

```bash
curl -X GET "http://localhost:4444/api/todos/?page=1&limit=10&priority=high&completed=false"
```

#### 3. Update Todo

```bash
curl -X PUT http://localhost:4444/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Go Advanced",
    "completed": true
  }'
```

#### 4. Hapus Todo

```bash
curl -X DELETE http://localhost:4444/api/todos/1
```

#### 5. Buat Kategori

```bash
curl -X POST http://localhost:4444/api/category/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Belajar",
    "color": "#FF5733"
  }'
```

---

## Database Skema

### Tabel Database

#### t_todos

Tabel utama untuk menyimpan data todo.

| Kolom       | Type      | Constraint                  | Deskripsi                   |
| ----------- | --------- | --------------------------- | --------------------------- |
| id          | INT       | PRIMARY KEY, AUTO_INCREMENT | Identitas unik              |
| title       | VARCHAR   | NOT NULL                    | Judul todo                  |
| description | TEXT      |                             | Deskripsi detail            |
| category_id | INT       | FOREIGN KEY                 | Referensi ke t_categories   |
| priority    | VARCHAR   |                             | Prioritas (high/medium/low) |
| completed   | BOOLEAN   | DEFAULT false               | Status penyelesaian         |
| due_date    | TIMESTAMP |                             | Tanggal deadline            |
| created_at  | TIMESTAMP | DEFAULT NOW()               | Waktu pembuatan             |

#### t_categories

Tabel untuk kategori todo.

| Kolom      | Type      | Constraint                  | Deskripsi        |
| ---------- | --------- | --------------------------- | ---------------- |
| id         | INT       | PRIMARY KEY, AUTO_INCREMENT | Identitas unik   |
| name       | VARCHAR   | NOT NULL                    | Nama kategori    |
| color      | VARCHAR   |                             | Kode warna (hex) |
| created_at | TIMESTAMP | DEFAULT NOW()               | Waktu pembuatan  |

### Relasi

**Hubungan One-to-Many:**

- Satu Category memiliki banyak Todo
- Foreign Key: `t_todos.category_id` → `t_categories.id`

### Model

**Todo Model (Go)**

```go
type Todo struct {
    ID          int       // Identitas unik
    Title       string    // Judul tugas
    Description string    // Deskripsi detail
    CategoryID  int       // Referensi kategori
    Category    Category  // Embedded category object
    Priority    string    // high, medium, low
    Completed   bool      // Status penyelesaian
    DueDate     time.Time // Tanggal deadline
    CreatedAt   *time.Time // Waktu pembuatan
}
```

**Category Model (Go)**

```go
type Category struct {
    ID        int       // Identitas unik
    Name      string    // Nama kategori
    Color     string    // Kode warna (hex)
    CreatedAt time.Time // Waktu pembuatan
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

[!erd](https://res.cloudinary.com/dmx8hcmxh/image/upload/v1778906893/todo.drawio_lrvdht.png)

Relasi: One-to-Many
- Satu kategori dapat memiliki banyak todo
- Setiap todo harus memiliki satu kategori
```

---
