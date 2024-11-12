package routes

import (

	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()

	ProviderRoutes(r)
	UserRoutes(r)
	TaskRoutes(r)
	SkillRoutes(r)

	return r
}