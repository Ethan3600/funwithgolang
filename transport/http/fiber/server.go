package http

import (
	"log"
	"strconv"

	"github.com/Ethan3600/funwithgolang/dtos"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"net/http"

	"github.com/Ethan3600/funwithgolang/application"
	cpu_controller "github.com/Ethan3600/funwithgolang/controllers/cpu"
)

func NewServer(app application.AppContext) {
	f := fiber.New()

	registerDefaultMiddleware(f)

	f.Get("/", func(c *fiber.Ctx) error {
		c.SendStatus(http.StatusOK)
		return c.SendString("Hello, World!")
	})

	registerV1Api(f.Group("/api/v1"), app)

	log.Fatal(f.Listen(":1323"))
}

func registerDefaultMiddleware(f *fiber.App) {
	f.Use(recover.New())
	f.Use(requestid.New())
	f.Use(logger.New())
}

func registerV1Api(v1Api fiber.Router, app application.AppContext) {
	v1Api.Get("/health", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.JSON(dtos.R{
			"status":        dtos.Success,
			"version":       app.Version,
			"database_type": app.C.Database.GetStrategy(),
		})
	})

	v1Api.Get("/cpu", func(c *fiber.Ctx) error {
		times, err := strconv.Atoi(c.Query("times", "1"))

		if err != nil {

			c.Status(http.StatusBadRequest)
			return c.JSON(dtos.R{
				"status":  dtos.Error,
				"message": "Invalid times Url",
			})
		}

		nums := cpu_controller.GetCpuIntensiveWork(times)

		c.Status(http.StatusOK)
		return c.JSON(dtos.R{
			"status": dtos.Success,
			"data":   nums,
		})
	})
}
