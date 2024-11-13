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
		tasks.PUT("/{id}", controllers.UpdateTask)
		tasks.POST("/{id}/accept", controllers.AcceptOffer)
		tasks.POST("/{id}/complete", controllers.AcceptTaskCompletion)
		// Add more task-related routes here
	}
}
