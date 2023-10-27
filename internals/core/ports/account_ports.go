package ports

import "github.com/gofiber/fiber/v2"

type AccountService interface {
	CreateAccount(firstName string, lastName string, password string) error
}

type AccountRepository interface {
	CreateAccount(firstName string, lastName string, password string) error
}

type AccountHandlers interface {
	CreateAccount(c *fiber.Ctx) error
}
