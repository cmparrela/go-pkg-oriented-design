package user

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
	return &handler{service}
}

// List godoc
// @Summary      List user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array} User
// @Router       /users [get]
func (h *handler) List(ctx *fiber.Ctx) error {
	users, err := h.service.List(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}

// Find godoc
// @Summary      Find user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200  {object} User
// @Router       /users/{id} [get]
func (h *handler) Find(ctx *fiber.Ctx) error {
	user, err := h.service.Find(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// Create godoc
// @Summary      Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} User
// @Router       /users [post]
func (h *handler) Create(ctx *fiber.Ctx) error {
	userDto := new(CreateDto)
	if err := ctx.BodyParser(userDto); err != nil {
		return err
	}

	user, err := h.service.Create(ctx.Context(), userDto)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(user)
}

// Update godoc
// @Summary      Update user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Param		 raw	body	object	true	"body raw"
// @Success      200  {object} User
// @Router       /users/{id} [put]
func (h *handler) Update(ctx *fiber.Ctx) error {
	userDto := new(UpdateDto)
	if err := ctx.BodyParser(userDto); err != nil {
		return err
	}

	result, err := h.service.Update(ctx.Context(), ctx.Params("id"), userDto)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(result)

}

// Delete godoc
// @Summary      Delete user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200
// @Router       /users/{id} [delete]
func (h *handler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)

}
