package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zekhoi/golang-todo/app/handlers"
	_ "github.com/zekhoi/golang-todo/app/middlewares"
)

func RegisterActivityRoute(router *gin.Engine, handlers handlers.ActivityHandler) {
	r := router.Group("/activity-groups")
	{
		r.GET("/", handlers.GetAllActivities)
		r.GET("/:id", handlers.GetActivity)
		r.POST("/", handlers.CreateActivity)
		r.PATCH("/:id", handlers.UpdateActivity)
		r.DELETE("/:id", handlers.DeleteActivity)
	}
}
