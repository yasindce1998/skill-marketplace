package models

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
	"golang.org/x/crypto/bcrypt"
)

// Provider model needs gorm.Model embedding for consistency
type Provider struct {
	gorm.Model            // Added gorm.Model for consistency
	Type          string    `json:"type" binding:"required,oneof=individual company"`
	FirstName     string    `json:"first_name" binding:"required"`
	LastName      string    `json:"last_name" binding:"required"`
	Email         string    `json:"email" binding:"required,email" gorm:"unique"`  // Added unique constraint
	Mobile        string    `json:"mobile" binding:"required"`
	Address       Address   `json:"address"`
	Company       Company   `json:"company"`
	Skills        []Skill   `json:"skills" gorm:"many2many:provider_skills;"` // Added skills relationship
	Tasks         []Task    `json:"tasks" gorm:"foreignKey:ProviderID"`       // Added tasks relationship
	Offers        []Offer   `json:"offers" gorm:"foreignKey:ProviderID"`      // Added offers relationship
}

type Address struct {
	gorm.Model          // Added gorm.Model
	ProviderID uint      `json:"-" gorm:"uniqueIndex"` // Changed to uniqueIndex
	StreetNo   string    `json:"street_no"`
	StreetName string    `json:"street_name"`
	City       string    `json:"city" binding:"required"`    // Added required validation
	Suburb     string    `json:"suburb"`
	State      string    `json:"state" binding:"required"`   // Added required validation
	PostCode   string    `json:"post_code" binding:"required"` // Added required validation
}

type Company struct {
	gorm.Model          // Added gorm.Model
	ProviderID     uint           `json:"-" gorm:"uniqueIndex"` // Changed to uniqueIndex
	Name           string         `json:"name" binding:"required"`
	Phone          string         `json:"phone" binding:"required"`
	TaxNumber      string         `json:"tax_number" binding:"required,alphanum,len=10"`
	Representative Representative `json:"representative"`
}

type Representative struct {
	gorm.Model          // Added gorm.Model
	CompanyID   uint      `json:"-" gorm:"uniqueIndex"` // Changed to uniqueIndex
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	Email       string    `json:"email" binding:"required,email" gorm:"unique"` // Added unique constraint
	Mobile      string    `json:"mobile" binding:"required"`
}

type Task struct {
	gorm.Model           // Added gorm.Model
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	Status      string         `json:"status" gorm:"default:'pending'"`
	ProviderID  sql.NullInt64  `json:"provider_id"`
	Provider    Provider       `json:"provider" gorm:"foreignKey:ProviderID"`
	DueDate     time.Time      `json:"due_date"`
	CompletedAt *time.Time     `json:"completed_at"`
	Offers      []Offer        `json:"offers" gorm:"foreignKey:TaskID"` // Added offers relationship
}
type Skill struct {
	gorm.Model           // Added gorm.Model
	Name            string         `json:"name" gorm:"unique;not null"`
	Description     string         `json:"description"`
	Category        string         `json:"category" gorm:"not null"`
	Level           int            `json:"level" gorm:"default:1"`
	IsActive        bool           `json:"is_active" gorm:"default:true"`
	ProviderID      sql.NullInt64  `json:"provider_id"`
	Provider        Provider       `json:"provider" gorm:"foreignKey:ProviderID"`
	Certification   string         `json:"certification"`
	YearsExperience int            `json:"years_experience" gorm:"default:0"`
	LastUsed        *time.Time     `json:"last_used"`
	CreatedBy       sql.NullInt64  `json:"created_by"`
	UpdatedBy       sql.NullInt64  `json:"updated_by"`
	Tags            []SkillTag     `json:"tags" gorm:"many2many:skill_tags;"`
	Users           []User         `json:"users" gorm:"many2many:user_skills;"` // Added users relationship
}
type SkillTag struct {
	gorm.Model           // Added gorm.Model
	Name        string    `json:"name" gorm:"unique;not null"`
	Description string    `json:"description"`
	Skills      []Skill   `json:"skills" gorm:"many2many:skill_tags;"`
}
type User struct {
	gorm.Model
	FirstName    string     `json:"first_name" gorm:"not null"`
	LastName     string     `json:"last_name" gorm:"not null"`
	Email        string     `json:"email" gorm:"unique;not null"`
	Password     string     `json:"-" gorm:"not null"` // "-" prevents password from being exposed in JSON
	Role         string     `json:"role" gorm:"default:'user'"`
	IsActive     bool       `json:"is_active" gorm:"default:true"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	PhoneNumber  string     `json:"phone_number"`
	Address      string     `json:"address"`
	ProfileImage string     `json:"profile_image"`
	Skills       []Skill    `json:"skills" gorm:"many2many:user_skills;"`
}

// Credentials represents the login credentials
type Credentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserResponse represents the safe user data to return in API responses
type UserResponse struct {
	ID          uint       `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	Role        string     `json:"role"`
	IsActive    bool       `json:"is_active"`
	LastLoginAt *time.Time `json:"last_login_at"`
}

// Offer represents an offer in the system
type Offer struct {
	gorm.Model
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Amount      float64    `json:"amount" gorm:"type:decimal(10,2);not null"`
	Currency    string     `json:"currency" gorm:"default:'USD'"`
	Status      string     `json:"status" gorm:"default:'pending'"`
	Type        string     `json:"type" gorm:"not null"` // hourly, fixed, project-based
	ValidUntil  *time.Time `json:"valid_until"`

	// Foreign keys and relationships
	ProviderID sql.NullInt64 `json:"provider_id"`
	Provider   Provider      `json:"provider" gorm:"foreignKey:ProviderID"`
	ClientID   uint          `json:"client_id" gorm:"not null"`
	Client     User          `json:"client" gorm:"foreignKey:ClientID"`
	TaskID     sql.NullInt64 `json:"task_id"`
	Task       Task          `json:"task" gorm:"foreignKey:TaskID"`

	// Additional fields
	Timeline       int         `json:"timeline" gorm:"comment:'Duration in days'"`
	Milestones     []Milestone `json:"milestones" gorm:"foreignKey:OfferID"`
	Requirements   string      `json:"requirements"`
	Terms          string      `json:"terms"`
	AttachmentURLs []string    `json:"attachment_urls" gorm:"type:text[]"`

	// Negotiation fields
	IsNegotiable bool     `json:"is_negotiable" gorm:"default:true"`
	MinAmount    *float64 `json:"min_amount" gorm:"type:decimal(10,2)"`
	MaxAmount    *float64 `json:"max_amount" gorm:"type:decimal(10,2)"`

	// Tracking fields
	AcceptedAt  *time.Time `json:"accepted_at"`
	RejectedAt  *time.Time `json:"rejected_at"`
	CompletedAt *time.Time `json:"completed_at"`
	CanceledAt  *time.Time `json:"canceled_at"`
}

// Milestone represents a milestone in an offer
type Milestone struct {
	gorm.Model
	OfferID     uint       `json:"offer_id" gorm:"not null"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Amount      float64    `json:"amount" gorm:"type:decimal(10,2);not null"`
	DueDate     time.Time  `json:"due_date"`
	Status      string     `json:"status" gorm:"default:'pending'"`
	CompletedAt *time.Time `json:"completed_at"`
}

// ValidateProvider validates provider data
func (p *Provider) ValidateProvider() error {
	if p.Email == "" {
		return errors.New("email is required")
	}
	if p.Type != "individual" && p.Type != "company" {
		return errors.New("invalid provider type")
	}
	return nil
}

// ValidateAddress validates address data
func (a *Address) ValidateAddress() error {
	if a.City == "" {
		return errors.New("city is required")
	}
	if a.State == "" {
		return errors.New("state is required")
	}
	if a.PostCode == "" {
		return errors.New("post code is required")
	}
	return nil
}

// ValidateTask validates task data
func (t *Task) ValidateTask() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.DueDate.Before(time.Now()) {
		return errors.New("due date must be in the future")
	}
	return nil
}

// BeforeCreate hook validates the offer before creation
func (o *Offer) BeforeCreate(tx *gorm.DB) error {
	if o.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if o.ValidUntil != nil && o.ValidUntil.Before(time.Now()) {
		return errors.New("valid until date must be in the future")
	}

	// Set default status if not provided
	if o.Status == "" {
		o.Status = "pending"
	}

	return nil
}

// ValidateOffer validates the offer data
func (o *Offer) ValidateOffer() error {
	if o.Title == "" {
		return errors.New("title is required")
	}

	if o.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if o.Type == "" {
		return errors.New("type is required")
	}

	if o.ClientID == 0 {
		return errors.New("client ID is required")
	}

	return nil
}

// Accept marks the offer as accepted
func (o *Offer) Accept(db *gorm.DB) error {
	now := time.Now()
	o.Status = "accepted"
	o.AcceptedAt = &now
	return db.Save(o).Error
}

// Reject marks the offer as rejected
func (o *Offer) Reject(db *gorm.DB) error {
	now := time.Now()
	o.Status = "rejected"
	o.RejectedAt = &now
	return db.Save(o).Error
}

// Complete marks the offer as completed
func (o *Offer) Complete(db *gorm.DB) error {
	now := time.Now()
	o.Status = "completed"
	o.CompletedAt = &now
	return db.Save(o).Error
}

// Cancel marks the offer as canceled
func (o *Offer) Cancel(db *gorm.DB) error {
	now := time.Now()
	o.Status = "canceled"
	o.CanceledAt = &now
	return db.Save(o).Error
}

// BeforeCreate hook to validate skill data before creation
func (s *Skill) BeforeCreate(tx *gorm.DB) error {
	if s.Level < 1 || s.Level > 5 {
		return errors.New("skill level must be between 1 and 5")
	}
	return nil
}

// ValidateSkill validates the skill data
func (s *Skill) ValidateSkill() error {
	if s.Name == "" {
		return errors.New("skill name is required")
	}
	if s.Category == "" {
		return errors.New("skill category is required")
	}
	if s.Level < 1 || s.Level > 5 {
		return errors.New("skill level must be between 1 and 5")
	}
	return nil
}

// BeforeCreate hook to hash password before saving to database
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Password == "" {
		return errors.New("password is required")
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ValidatePassword checks if the provided password matches the stored hash
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		Role:        u.Role,
		IsActive:    u.IsActive,
		LastLoginAt: u.LastLoginAt,
	}
}

// UpdateLastLogin updates the user's last login timestamp
func (u *User) UpdateLastLogin(db *gorm.DB) error {
	now := time.Now()
	u.LastLoginAt = &now
	return db.Save(u).Error
}
