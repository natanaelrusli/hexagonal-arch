package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natanaelrusli/hexagonal-arch/internals/core/ports"
)

type UserHandlers struct {
	userService ports.UserService
}

var _ ports.UserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) Login(c *fiber.Ctx) error {
	var email string
	var password string

	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserHandlers) Register(c *fiber.Ctx) error {
	var email string
	var password string
	var confirmPassword string

	err := h.userService.Register(email, password, confirmPassword)
	if err != nil {
		return err
	}
	return nil
}
