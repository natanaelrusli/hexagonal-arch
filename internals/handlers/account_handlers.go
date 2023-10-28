package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natanaelrusli/hexagonal-arch/internals/core/ports"
	"github.com/natanaelrusli/hexagonal-arch/internals/dto/requests"
)

type AccountHandlers struct {
	accountService ports.AccountService
}

func NewAccountHandlers(accountService ports.AccountService) *AccountHandlers {
	return &AccountHandlers{
		accountService: accountService,
	}
}

func (h *AccountHandlers) CreateAccount(c *fiber.Ctx) error {
	req := new(requests.CreateAccountRequest)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	err := h.accountService.CreateAccount(req.FirstName, req.LastName, req.Password)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "new account created",
	})
}
