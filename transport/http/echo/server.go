package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Ethan3600/funwithgolang/application"
	cpu_controller "github.com/Ethan3600/funwithgolang/controllers/cpu"
	person_controller "github.com/Ethan3600/funwithgolang/controllers/person"
	"github.com/Ethan3600/funwithgolang/dtos"
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
			"database_type": app.C.Database.GetStrategy(),
		})
	})

	people := v1Api.Group("/people")
	registerPeopleApi(*people, app)

	v1Api.GET("/cpu", func(c echo.Context) error {
		times := int(1)

		err := echo.QueryParamsBinder(c).
			Int("times", &times).
			BindError()

		if err != nil {
			return c.JSON(http.StatusBadRequest, dtos.R{
				"status":  dtos.Error,
				"message": "Invalid times Url",
			})
		}

		nums := cpu_controller.GetCpuIntensiveWork(times)
		return c.JSON(http.StatusOK, dtos.R{
			"status": dtos.Success,
			"data":   nums,
		})
	})
}

func registerPeopleApi(people echo.Group, app application.AppContext) {
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
				"message": fmt.Sprintf("Person with ID: %s not found", id),
			})
		}

		return c.JSON(http.StatusOK, dtos.R{
			"status": dtos.Success,
			"data":   person,
		})
	})
}
