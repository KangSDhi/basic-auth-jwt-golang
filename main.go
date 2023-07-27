package main

import (
	"basic-jwt-auth/config"
	"basic-jwt-auth/handlers"
	"basic-jwt-auth/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	app.Post("/login", handlers.Login)
	app.Get("/protected", jwt, handlers.Protected)
	app.Listen(":3000")
}
