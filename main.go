package main

import (
	"go-fiber/databases/conn"
	"go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conn.ConnDB()

	app := fiber.New()

	routes.Routes(app)

	app.Listen("127.0.0.1:8000")
}