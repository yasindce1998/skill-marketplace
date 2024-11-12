package controllers

import (
	"github.com/yasindce1998/skill-marketplace/api/models"
	"github.com/yasindce1998/skill-marketplace/config"
	"net/http"

	"github.com/gin-gonic/gin"
)
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}