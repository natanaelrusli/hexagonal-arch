package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/natanaelrusli/hexagonal-arch/internals/core/ports"
)

// Server represents the main server structure.
type Server struct {
	app *fiber.App
}

// NewServer creates a new server instance.
func NewServer() *Server {
	return &Server{
		app: fiber.New(),
	}
}

// InitializeUserHandlers initializes user-related handlers.
func (s *Server) InitializeUserHandlers(uHandlers ports.UserHandlers) {
	v1 := s.app.Group("/v1/user")
	v1.Post("/login", uHandlers.Login)
	v1.Post("/register", uHandlers.Register)
}

// InitializeAccountHandlers initializes account-related handlers.
func (s *Server) InitializeAccountHandlers(accHandlers ports.AccountHandlers) {
	v1 := s.app.Group("/v1/account")
	v1.Post("/create", accHandlers.CreateAccount)
	// Add other account handlers as needed...
}

// StartServer starts the server on the specified port.
func (s *Server) StartServer(port string) {
	err := s.app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
