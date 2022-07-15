package book

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	List(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}

// List godoc
// @Summary      List book
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array} Book
// @Router       /books [get]
func (b *handler) List(ctx *fiber.Ctx) error {
	books, err := b.service.List(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(books)
}

// Find godoc
// @Summary      Find book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200  {object} Book
// @Router       /books/{id} [get]
func (b *handler) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	book, err := b.service.Find(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(book)
}

// Create godoc
// @Summary      Create book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} Book
// @Router       /books [post]
func (b *handler) Create(ctx *fiber.Ctx) error {
	bookDto := new(CreateDto)
	if err := ctx.BodyParser(bookDto); err != nil {
		return err
	}

	book, err := b.service.Create(ctx.Context(), bookDto)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(book)
}

// Update godoc
// @Summary      Update book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Param		 raw	body	object	true	"body raw"
// @Success      200  {object} Book
// @Router       /books/{id} [put]
func (b *handler) Update(ctx *fiber.Ctx) error {
	bookDto := new(UpdateDto)
	if err := ctx.BodyParser(bookDto); err != nil {
		return err
	}

	id := ctx.Params("id")

	result, err := b.service.Update(ctx.Context(), id, bookDto)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(result)

}

// Delete godoc
// @Summary      Delete book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200
// @Router       /books/{id} [delete]
func (b *handler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := b.service.Delete(ctx.Context(), id); err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)

}
