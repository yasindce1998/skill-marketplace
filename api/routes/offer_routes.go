package routes

import (
	"github.com/yasindce1998/skill-marketplace/api/controllers"

	"github.com/gin-gonic/gin"
)

func OfferRoutes(r *gin.Engine) {
	offers := r.Group("/offers")
	{
		offers.POST("", controllers.CreateOffer)
	}
}