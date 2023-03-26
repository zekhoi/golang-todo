package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/zekhoi/golang-todo/app/dto"
	_ "github.com/zekhoi/golang-todo/app/entities"
	_ "github.com/zekhoi/golang-todo/app/repositories"
	"github.com/zekhoi/golang-todo/app/services"
)

type ToDoHandler interface {
	GetAllToDos(c *gin.Context)
	GetToDo(c *gin.Context)
	CreateToDo(c *gin.Context)
	UpdateToDo(c *gin.Context)
	DeleteToDo(c *gin.Context)
}

type toDoHandler struct {
	service services.ToDoService
}

func NewToDoHandler(service services.ToDoService) ToDoHandler {
	return &toDoHandler{service: service}
}

func (h *toDoHandler) GetAllToDos(c *gin.Context) {
	// ...
}

func (h *toDoHandler) GetToDo(c *gin.Context) {
	// ...
}

func (h *toDoHandler) CreateToDo(c *gin.Context) {
	// ...
}

func (h *toDoHandler) UpdateToDo(c *gin.Context) {
	// ...
}

func (h *toDoHandler) DeleteToDo(c *gin.Context) {
	// ...
}
