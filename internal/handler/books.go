package handler

import (
	"elasticSearch/internal/app/constants"
	"elasticSearch/internal/handler/response"
	"elasticSearch/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateBook(c *fiber.Ctx) error {
	var input models.Book
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}

	id, err := h.service.Books.CreateBook(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(response.NewResponseMessage(fiber.StatusOK, id))
}

func (h *Handler) Search(c *fiber.Ctx) error {
	searchInput := c.Query(constants.SearchParam)

	books, err := h.service.Books.Search(searchInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func (h *Handler) DeleteById(c *fiber.Ctx) error {
	var ids models.DeleteIds
	if err := c.BodyParser(&ids); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}

	if err := h.service.Delete(ids); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(response.NewResponseMessage(fiber.StatusOK, constants.Success))
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var input models.Book
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}
	if input.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}
	if err := h.service.Update(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(response.NewResponseMessage(fiber.StatusOK, constants.Success))
}

func (h *Handler) GetBook(c *fiber.Ctx) error {
	searchInput := c.Query(constants.SearchParam)

	book, err := h.service.Books.GetFormCache(searchInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *Handler) Sync(c *fiber.Ctx) error {
	if err := h.service.Books.Sync(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response.NewResponseMessage(fiber.StatusOK, constants.SuccessSync))
}
