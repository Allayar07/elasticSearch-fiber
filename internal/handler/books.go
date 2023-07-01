package handler

import (
	"elasticSearch/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Hello(c *fiber.Ctx) error {
	_, err := c.WriteString("Hello World")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}

	return nil
}

func (h *Handler) CreateBook(c *fiber.Ctx) error {
	var input models.Book
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	id, err := h.service.Books.CreateBook(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (h *Handler) Search(c *fiber.Ctx) error {
	searchInput := c.Query("find")

	books, err := h.service.Books.Search(searchInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func (h *Handler) deleteById(c *fiber.Ctx) error {
	var ids models.DeleteIds
	if err := c.BodyParser(&ids); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.service.Delete(ids); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "OK",
	})
}
func (h *Handler) Update(c *fiber.Ctx) error {
	var input models.Book
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.service.Update(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "OK",
	})
}