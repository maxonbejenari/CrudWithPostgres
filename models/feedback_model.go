package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type Feedback struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Name      string    `gorm:"not null" json:"name,omitempty"`
	Email     string    `gorm:"not null" json:"email,omitempty"`
	Feedback  string    `gorm:"uniqueIndex:idx_feedback;not null" json:"feedback,omitempty"`
	Rating    *float32  `gorm:"not null" json:"rating,omitempty"`
	Status    string    `json:"status,omitempty"`
	CreatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
}

func (feedback *Feedback) BeforeCreate() (err error) {
	// uuid (Universal Unique Identifier)
	feedback.ID = uuid.New().String() //create unique ID
	return nil
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

// this func verify if my model respect any requirments about my model "Feedback"
func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	errorValidatorType := err.(validator.ValidationErrors)
	if err != nil {
		for _, err := range errorValidatorType {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type CreateFeedbackSchema struct {
	Name     string   `json:"name" validate:"required"`
	Email    string   `json:"email" validate:"required"`
	Feedback string   `json:"feedback" validate:"required"`
	Rating   *float32 `json:"rating" validate:"required"`
	Status   string   `json:"status,omitempty"`
}

type UpdateFeedbackSchema struct {
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email,omitempty"`
	Feedback string   `json:"feedback,omitempty"`
	Rating   *float32 `json:"rating,omitempty"`
	Status   string   `json:"status,omitempty"`
}
