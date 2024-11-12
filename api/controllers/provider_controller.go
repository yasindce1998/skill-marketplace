package controllers

import (
	"github.com/yasindce1998/skill-marketplace/api/models"
	"github.com/yasindce1998/skill-marketplace/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProvider godoc
// @Summary Create a new provider
// @Description Create a new provider in the system
// @Tags providers
// @Accept json
// @Produce json
// @Param provider body models.Provider true "Provider information"
// @Success 201 {object} models.Provider
// @Failure 400 {object} gin.H
// @Router /providers [post]
func CreateProvider(c *gin.Context) {
	var provider models.Provider
	if err := c.ShouldBindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&provider).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, provider)
}

func GetProviders(c *gin.Context) {
	var providers []models.Provider
	if err := config.DB.Find(&providers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, providers)
}

func UpdateProvider(c *gin.Context) {
	var provider models.Provider
	if err := c.ShouldBindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&provider).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, provider)
}
