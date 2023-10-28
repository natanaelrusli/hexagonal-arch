package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natanaelrusli/hexagonal-arch/internals/core/domain"
)

type AccountService interface {
	CreateAccount(firstName string, lastName string, password string) error
}

type AccountRepository interface {
	CreateAccount(*domain.Account) error
}

type AccountHandlers interface {
	CreateAccount(c *fiber.Ctx) error
}
