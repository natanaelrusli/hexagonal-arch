package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natanaelrusli/hexagonal-arch/internals/core/ports"
	requests "github.com/natanaelrusli/hexagonal-arch/internals/dto"
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
	req := new(requests.UserRegisterRequest)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	err := h.userService.Register(req.Email, req.Password, req.ConfirmPassword)

	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "new user registered",
	})
}
