package server

import (
	"backend-coding-challenge/config"
	"backend-coding-challenge/database"
	"backend-coding-challenge/handler"
	"backend-coding-challenge/migration"
	"backend-coding-challenge/repository"
	"backend-coding-challenge/router"
	"backend-coding-challenge/service"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Config *config.Config
	Router *gin.Engine
	DB *gorm.DB	
}

func New(configPath string) *Server {
	// Load config
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Init Postgres
	db, err := database.InitPostgres(cfg.App.Db.Postgres)
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}
	fmt.Println("postgres connected")

	// Gin engine
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	fmt.Println("cors initialized")

	srv := &Server{
		Config: cfg,
		DB:     db,
		Router: r,
	}

	return srv
}

// init new modules here
func (s *Server) Todo() {
	api := s.Router.Group("/api/todos/")

	repo := repository.NewTodoRepository(s.DB)
	svc := service.NewTodoService(repo)
	categoryRepo := repository.NewCategoryRepository(s.DB)
	categorySvc := service.NewCategoryService(categoryRepo)
	handler := handler.NewTodoHandler(svc, categorySvc)

	r := router.NewTodoRouter(handler)
	r.InitTodoRoutes(api)
}

func (s *Server) Category() {
	api := s.Router.Group("api/category/")

	repo := repository.NewCategoryRepository(s.DB)
	svc := service.NewCategoryService(repo)
	handler := handler.NewCategoryHandler(svc)

	r := router.NewCategoryRouter(handler)
	r.InitCategoryRoutes(api)
}

func (s *Server) Migrate() {
	migrationConfig := s.Config.Migration
	if migrationConfig.Enabled {
		fmt.Println("=== Running Manual Migrations ===")
		if err := migration.RunAll(s.DB); err != nil {
			log.Fatalf("migration failed: %v", err)
		}

		// seed initial data
		fmt.Println("\n=== Seeding Initial Data ===")
		database.SeedCategories(s.DB)
		database.SeedTodos(s.DB)
		fmt.Println("✓ Data seeding completed")
	}
}

func (s *Server) Run() {
	port := s.Config.App.Port
	fmt.Printf("server running on %s\n", port)
	if err := s.Router.Run(port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}



