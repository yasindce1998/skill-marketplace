package routes

import (
	"github.com/yasindce1998/skill-marketplace/api/controllers"

	"github.com/gin-gonic/gin"
)

func SkillRoutes(r *gin.Engine) {
	skills := r.Group("/skills")
	{
		skills.POST("", controllers.CreateSkill)
		skills.PUT("", controllers.UpdateSkill)
		// Add more skill-related routes here
	}
}