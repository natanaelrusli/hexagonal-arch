package main

import (
	"github.com/natanaelrusli/hexagonal-arch/internals/core/services"
	"github.com/natanaelrusli/hexagonal-arch/internals/handlers"
	"github.com/natanaelrusli/hexagonal-arch/internals/repositories"
	"github.com/natanaelrusli/hexagonal-arch/internals/server"
)

// MongoConn:string > UserRepository > UserService > UserHandler > Server

func main() {
	mongoConn := "secrett"

	userRepository, _ := repositories.NewUserRepository(mongoConn)
	userService := services.NewUserService(userRepository)
	userHandlers := handlers.NewUserHandlers(userService)

	httpServer := server.NewServer(
		userHandlers,
	)

	httpServer.Initialize()
}
