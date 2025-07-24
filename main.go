package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"todo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, but continuing...")
		}
	}

	// Fetch environment variables
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

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
		log.Fatal("Internal server error", err)
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

	app.Listen(":8080")

}
