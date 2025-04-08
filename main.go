package main

import (
	"CRUDwPOSTGRES/controller"
	"CRUDwPOSTGRES/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	env, err := initializers.LoadEnv(".")
	if err != nil {
		log.Fatal("Could not load env vars", err)
	}
	initializers.ConnectDB(&env)
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)

	//app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Context-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	micro.Route("/feedbacks", func(router fiber.Router) {
		router.Post("/create", controller.CreateFeedbackHandler)
		router.Get("/find", controller.FindFeedbackHandler)
	})

	micro.Route("/feedbacks/:feedbackId", func(router fiber.Router) {
		router.Get("", controller.FindFeedbackByIdHandler)
		router.Patch("", controller.UpdateFeedbackHandler)
		router.Delete("", controller.DeleteFeedbackHandler)
	})

	micro.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "OK",
		})
	})

	log.Fatal(app.Listen(":" + env.ServerPort))
}
