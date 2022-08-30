package http

import (
	"strconv"

	"github.com/Ethan3600/funwithgolang/dtos"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/requestid"

	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/Ethan3600/funwithgolang/application"
	cpu_controller "github.com/Ethan3600/funwithgolang/controllers/cpu"
)

func NewServer(app application.AppContext) {
	g := gin.Default()
	pprof.Register(g)

	registerDefaultMiddleware(g)

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	registerV1Api(g.Group("/api/v1"), app)

	g.Run(":1323")
}

func registerDefaultMiddleware(g *gin.Engine) {
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(requestid.New())
}

func registerV1Api(v1Api *gin.RouterGroup, app application.AppContext) {
	v1Api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, dtos.R{
			"status":        dtos.Success,
			"version":       app.Version,
			"database_type": app.C.Database.GetStrategy(),
		})
	})

	v1Api.GET("/cpu", func(c *gin.Context) {
		times, err := strconv.Atoi(c.DefaultQuery("times", "1"))

		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.R{
				"status":  dtos.Error,
				"message": "Invalid times Url",
			})
		}

		nums := cpu_controller.GetCpuIntensiveWork(times)
		c.JSON(http.StatusOK, dtos.R{
			"status": dtos.Success,
			"data":   nums,
		})
	})
}
