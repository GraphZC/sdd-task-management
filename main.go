package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GraphZC/sdd-task-management/configs"
	"github.com/GraphZC/sdd-task-management/domain/usecases"
	"github.com/GraphZC/sdd-task-management/internal/adapters/mysql"
	"github.com/GraphZC/sdd-task-management/internal/adapters/rest"
	"github.com/GraphZC/sdd-task-management/middlewares"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func main() {
	app := fiber.New()

	ctx := context.Background()

	cfg := configs.NewConfig()

	db, err := sqlx.ConnectContext(ctx, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := mysql.NewUserMySQLRepository(db)
	userService := usecases.NewUserService(userRepo, cfg)
	userHandler := rest.NewUserHandler(userService)

	taskRepo := mysql.NewTaskMySQLRepository(db)
	taskService := usecases.NewTaskService(taskRepo)
	taskHandler := rest.NewTaskHandler(taskService)

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	app.Use(middlewares.JwtMiddleware(cfg.JWTSecret))
	app.Post("/task", taskHandler.CreateTask)
	app.Get("/task", taskHandler.FindTaskByUserID)
	app.Get("/task/:taskID", taskHandler.FindTaskByID)
	app.Delete("/task/:taskID", taskHandler.DeleteTaskByID)
	app.Put("/task/:taskID", taskHandler.UpdateTaskByID)
	app.Post("/task/:taskID/status", taskHandler.UpdateTaskStatusByID)

	if err := app.Listen(":9000"); err != nil {
		log.Fatal(err)
	}
}
