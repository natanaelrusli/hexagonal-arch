package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/natanaelrusli/hexagonal-arch/internals/core/ports"
)

type Server struct {
	// put every new handlers here
	userHandlers ports.UserHandlers
	// middlewares ports.IMiddlewares
	// paymentHandlers ports.IPaymentHandlers
}

func NewServer(uHandlers ports.UserHandlers) *Server {
	return &Server{
		userHandlers: uHandlers,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("/user")
	userRoutes.Post("/login", s.userHandlers.Login)
	userRoutes.Post("/register", s.userHandlers.Register)

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
