package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echoapp/application"
	person_controller "echoapp/controllers/person"
	"echoapp/dtos"
)

func NewServer(app application.AppContext) {
	e := echo.New()

	registerDefaultMiddleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	registerV1Api(e.Group("/api/v1"), app)

	e.Logger.Fatal(e.Start(":1323"))
}

func registerDefaultMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
}

func registerV1Api(v1Api *echo.Group, app application.AppContext) {

	v1Api.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dtos.R{
			"status":        dtos.Success,
			"version":       app.Version,
			"database_type": app.Db.GetStrategy(),
		})
	})

	people := v1Api.Group("/people")
	people.POST("", func(c echo.Context) error {
		var person dtos.Person

		if err := c.Bind(&person); err != nil {
			c.Logger().Fatal(err)
			return c.JSON(http.StatusBadRequest, dtos.R{
				"status":  dtos.Error,
				"message": "Unable to bind request payload",
			})
		}

		new_id, err := person_controller.CreatePerson(person, app)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dtos.R{
				"status":  dtos.Error,
				"message": "Failed to save person",
			})
		}

		return c.JSON(http.StatusCreated, dtos.R{
			"status": dtos.Success,
			"id":     new_id,
		})
	})

	people.GET("", func(c echo.Context) error {
		people, err := person_controller.GetPeople(app)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dtos.R{
				"status":  dtos.Error,
				"message": "Failed to get people",
			})
		}

		return c.JSON(http.StatusOK, dtos.R{
			"status": dtos.Success,
			"total":  len(people),
			"data":   people,
		})
	})

	people.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")

		person, err := person_controller.GetPerson(id, app)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dtos.R{
				"status":  dtos.Error,
				"message": "Internal failure",
			})
		}

		if person == nil {
			return c.JSON(http.StatusNotFound, dtos.R{
				"status":  dtos.Fail,
				"message": "Internal failure",
			})
		}

		return c.JSON(http.StatusOK, dtos.R{
			"status": dtos.Success,
			"data":   person,
		})
	})
}