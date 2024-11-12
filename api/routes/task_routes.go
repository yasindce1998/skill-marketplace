package routes

import (
	"github.com/yasindce1998/skill-marketplace/api/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	tasks := r.Group("/tasks")
	{
		tasks.POST("", controllers.CreateTask)
		tasks.PUT("/progress", controllers.UpdateTaskProgress)
		// Add more task-related routes here
	}
}
