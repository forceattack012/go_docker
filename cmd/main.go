package main

import (
	"log"
	"os"

	"github.com/forceattack012/go_docker/todo"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("docker.env")
	if err != nil {
		log.Fatalf("Could not load env %s", err.Error())
	}
}

func main() {
	dsn := os.Getenv("dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(todo.Todo{})

	handler := todo.NewTodoHandler(db)
	r := gin.Default()

	r.GET("/", handler.SayHello)
	r.POST("/", handler.Receiver)
	r.POST("/todo", handler.CreateTodo)
	r.GET("/todo", handler.GetTodo)

	r.Run()
}
