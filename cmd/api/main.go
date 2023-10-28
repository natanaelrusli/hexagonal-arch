package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/natanaelrusli/hexagonal-arch/internals/core/services"
	"github.com/natanaelrusli/hexagonal-arch/internals/handlers"
	"github.com/natanaelrusli/hexagonal-arch/internals/repositories"
	"github.com/natanaelrusli/hexagonal-arch/internals/server"
)

// MongoConn:string > UserRepository > UserService > UserHandler > Server

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	loadEnv()
	mongoConn := "mongodb+srv://admin:admin@cluster0.dmvyscs.mongodb.net/?retryWrites=true&w=majority"

	_, err := repositories.NewPostgresAccountRepository()

	if err != nil {
		log.Fatal(err)
	}

	userRepository, _ := repositories.NewUserRepository(mongoConn)
	userService := services.NewUserService(userRepository)
	userHandlers := handlers.NewUserHandlers(userService)

	accountRepository, _ := repositories.NewPostgresAccountRepository()
	accountService := services.NewAccountService(accountRepository)
	accountHandlers := handlers.NewAccountHandlers(accountService)

	server := server.NewServer()
	server.InitializeAccountHandlers(accountHandlers)
	server.InitializeUserHandlers(userHandlers)

	server.StartServer("8000")
}
