package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"todo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_PORT     = os.Getenv("DB_PORT")
)

func main() {

	// Create an instance of the application
	app := fiber.New(fiber.Config{
		Prefork: false,
		AppName: "Todo_Go",
	})

	app.Use(cors.New())
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Nairobi",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Internal server error")
	}

	sqlDB, _ := db.DB()

	//migrate the schema whenever possible
	db.AutoMigrate(&todo.Todo{})

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Instantiate controllers
	controller := todo.TodoController{
		Db: db,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Post("api/v1/todo", controller.CreateTodo)
	app.Get("api/v1/todo", controller.GetAllToDo)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
