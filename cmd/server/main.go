package main
import (
  "github.com/gofiber/fiber/v3"
  "log"
  "time"
)

var (
  GIT_COMMIT = ""
  IMAGE_NAME = ""
)

func main() {
  // Show provenance information
  log.Printf("APP GIT_COMMIT = %s", GIT_COMMIT)
  log.Printf("APP IMAGE_NAME = %s", IMAGE_NAME)

  // Init fiber app
  app := fiber.New()

  app.Get("/about", func(c fiber.Ctx) error {
    return c.JSON(fiber.Map{ "GIT_COMMIT": GIT_COMMIT, "IMAGE_NAME": IMAGE_NAME })
  })

  // Define route
  app.Get("/", func(c fiber.Ctx) error {
    return c.JSON(fiber.Map{ "message": "My name is 0hana", "timestamp": time.Now().UnixMilli() })
    // return c.JSON(fiber.Map{ "message": "My name is 0hana", "timestamp": time.Now().Format("01021504")})
    // Go uses an interesting "Reference Date" string to specify format,
    // with a 1, 2, 3, 4, 5, 6 pattern:
    //
    // 1. Month
    // 2. Day
    // 3. Hour (12+3=15 for 24hr time)
    // 4. Minute
    // 5. Second
    // 6. Year
    //
    // So "01021504" means:
    //             a 2-digit month  (01-12),
    // followed by a 2-digit day    (01-31),
    // followed by a 2-digit hour   (24hr vs 12hr am/pm),
    // followed by a 2-digit minute (00-59)
    //
    // Whereas "04150201" is the reverse:
    //             a 2-digit minute (00-59),
    // followed by a 2-digit hour   (24hr vs 12hr am/pm),
    // followed by a 2-digit day    (01-31),
    // followed by a 2-digit month  (01-12)
  })

  // Start server on port 80
  log.Fatal(app.Listen(":80"))
}
