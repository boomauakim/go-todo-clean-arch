package main

import (
	todoRouter "github.com/boomauakim/go-todo-clean-arch/todo/delivery/http/route"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	todoRouter.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))
}
