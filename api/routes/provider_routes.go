package routes

import (
	"github.com/yasindce1998/skill-marketplace/api/controllers"

	"github.com/gin-gonic/gin"
)

func ProviderRoutes(r *gin.Engine) {
	providers := r.Group("/providers")
	{
		providers.POST("", controllers.CreateProvider)
		// Add more provider-related routes here
	}
}