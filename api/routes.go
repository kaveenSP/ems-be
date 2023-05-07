package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EndpointManager(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	g := e.Group("/ems-planners/api")
	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	g.POST("/teacher", CreateTeacherApi)
	g.GET("/teacher", FindTeacherApi)
	g.GET("/teachers", FindAllTeachersApi)
	g.PUT("/teacher", UpdateTeacherApi)
	g.DELETE("/teacher", DeleteTeacherApi)
	g.POST("/event", CreateEventApi)
	g.GET("/event", FindEventApi)
	g.GET("/events", FindAllEventsApi)
	g.PUT("/event", UpdateEventApi)
	g.DELETE("/event", DeleteEventApi)
}
