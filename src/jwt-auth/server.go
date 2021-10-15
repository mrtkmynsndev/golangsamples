package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	Port string
}

type Server struct {
	config ServerConfig
}

func New(config ServerConfig) *Server {
	return &Server{
		config: config,
	}
}

func Start() {
	app := fiber.New(fiber.Config{})

	app.Post("/signup", SignUp)

	app.Post("/signin", SignIn)

	app.Use("/users", AuthenticationMiddleware())
	app.Get("/users", UserListHandler)

	app.Use("/user", AuthenticationMiddleware())
	app.Get("/user", GetUserHandler)

	log.Fatal(app.Listen(":3000"))
}
