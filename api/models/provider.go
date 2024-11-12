package models

import (
	"time"
)

type Provider struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Type      string    `json:"type" binding:"required,oneof=individual company"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Mobile    string    `json:"mobile" binding:"required"`
	Address   Address   `json:"address"`
	Company   Company   `json:"company"`
}

type Address struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ProviderID uint   `gorm:"index" json:"-"`
	StreetNo   string `json:"street_no"`
	StreetName string `json:"street_name"`
	City       string `json:"city"`
	Suburb     string `json:"suburb"`
	State      string `json:"state"`
	PostCode   string `json:"post_code"`
}

type Company struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	ProviderID     uint           `gorm:"index" json:"-"`
	Name           string         `json:"name" binding:"required"`
	Phone          string         `json:"phone" binding:"required"`
	TaxNumber      string         `json:"tax_number" binding:"required,alphanum,len=10"`
	Representative Representative `json:"representative"`
}

type Representative struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	CompanyID uint   `gorm:"index" json:"-"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Mobile    string `json:"mobile" binding:"required"`
}
