package controllers

import (
	"github.com/yasindce1998/skill-marketplace/api/models"
	"github.com/yasindce1998/skill-marketplace/config"
	"time"
)

type PeriodicStats struct {
	TotalTasks             int     `json:"total_tasks"`
	CompletedTasks         int     `json:"completed_tasks"`
	RejectedTasks          int     `json:"rejected_tasks"`
	AvgProviderSuccessRatio float64 `json:"avg_provider_success_ratio"`
}

func GetProviderCount() int {
	var providers []models.Provider
	config.DB.Find(&providers)
	return len(providers)
}

func GetPeriodicStats(startDate, endDate time.Time) PeriodicStats {
	var tasks []models.Task
	config.DB.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&tasks)

	var completedTasks, rejectedTasks int
	var providerSuccessRatios []float64

	for _, task := range tasks {
		if task.Status == "completed" {
			completedTasks++
		} else if task.Status == "rejected" {
			rejectedTasks++
		}

		if task.ProviderID.Valid {
			var provider models.Provider
			config.DB.First(&provider, task.ProviderID.Int64)
			providerSuccessRatios = append(providerSuccessRatios, float64(completedTasks)/(float64(completedTasks+rejectedTasks))*100)
		}
	}

	var avgProviderSuccessRatio float64
	if len(providerSuccessRatios) > 0 {
		for _, ratio := range providerSuccessRatios {
			avgProviderSuccessRatio += ratio
		}
		avgProviderSuccessRatio /= float64(len(providerSuccessRatios))
	}

	return PeriodicStats{
		TotalTasks:             len(tasks),
		CompletedTasks:         completedTasks,
		RejectedTasks:          rejectedTasks,
		AvgProviderSuccessRatio: avgProviderSuccessRatio,
	}
}