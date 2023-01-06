package main

import (
	"github.com/forceattack012/example_docker/todo"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=todopassword dbname=postgres port=5439 sslmode=disable TimeZone=Asia/Bangkok"
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
