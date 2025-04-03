package controller

import (
	"CRUDwPOSTGRES/initializers"
	"CRUDwPOSTGRES/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

func CreateFeedbackHandler(c *fiber.Ctx) error {
	var payload *models.CreateFeedbackSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	now := time.Now()
	newFeedback := models.Feedback{
		Name:      payload.Name,
		Email:     payload.Email,
		Feedback:  payload.Feedback,
		Rating:    payload.Rating,
		Status:    payload.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := initializers.DB.Create(&newFeedback)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "Error Failed",
			"message": "Feedback already exists",
		})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"note": newFeedback},
	})
}

func FindFeedbackHandler(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var feedbacks []models.Feedback
	result := initializers.DB.Limit(intLimit).Offset(offset).Find(&feedbacks)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": result.Error,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":    "Success",
		"result":    len(feedbacks),
		"feedbacks": feedbacks,
	})
}
