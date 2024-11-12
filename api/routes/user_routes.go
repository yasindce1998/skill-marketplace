package routes

import (
	"github.com/yasindce1998/skill-marketplace/api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", controllers.CreateUser)
		// Add more user-related routes here
	}
}