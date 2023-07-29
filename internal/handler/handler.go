package handler

import (
	"elasticSearch/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(logger.New())

	app.Post("/create", h.CreateBook)
	app.Get("/search", h.Search)
	app.Delete("/delete", h.deleteById)
	app.Put("/update", h.Update)
	app.Get("/get", h.GetBook)
	app.Post("/sync", h.Sync)

	return app

}
