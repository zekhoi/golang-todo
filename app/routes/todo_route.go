package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zekhoi/golang-todo/app/handlers"
)

func RegisterToDoRoute(router *gin.Engine, handlers handlers.ToDoHandler) {
	r := router.Group("/todo-items")
	{
		r.GET("/", handlers.GetAllToDos)
		r.GET("/:id", handlers.GetToDo)
		r.POST("/", handlers.CreateToDo)
		r.PATCH("/:id", handlers.UpdateToDo)
		r.DELETE("/:id", handlers.DeleteToDo)
	}
}
