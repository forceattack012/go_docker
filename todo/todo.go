package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string `json:"title"`
}

type TodoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

func (h *TodoHandler) SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world !!",
	})
}

func (h *TodoHandler) Receiver(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := h.db.Create(&todo)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": todo.ID,
	})
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	var todo []Todo
	result := h.db.Find(&todo)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}
