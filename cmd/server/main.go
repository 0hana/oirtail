package main
import (
  "log"
  "github.com/gofiber/fiber/v3"
)

func main() {
  // Init fiber app
  app := fiber.New()

  // Define route
  app.Get("/", func(c fiber.Ctx) error {
    return c.SendString("Running Fiber v3")
  })

  // Start server on port 8080
  log.Fatal(app.Listen(":8080"))
}
