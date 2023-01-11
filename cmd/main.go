package main

import (
	"fmt"
	"log"

	"github.com/forceattack012/go_docker/todo"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
	Sslmode  string `yaml:"sslmode"`
	Timezone string `yaml:"timezone"`
	Env      string `yaml:"env"`
	App_Port string `yaml:"app_port"`
}

var cfg Config

func init() {
	cfg.ReadConfig("./config")
	fmt.Printf("Env : %s \n", cfg.Env)
}

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Dbname, cfg.Port, cfg.Sslmode, cfg.Timezone)
	fmt.Printf("%s\n", dsn)
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

	r.Run(":" + cfg.App_Port)
}

func (c *Config) ReadConfig(file string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(file)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf(err.Error())
	}
}
