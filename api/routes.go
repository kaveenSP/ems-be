package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EndpointManager(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowCredentials: true,
	}))

	g := e.Group("/ems-planners/api")
	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	g.POST("/auth/login", LoginApi)
	g.POST("/admin", CreateAdminApi)
	g.PUT("/admin", UpdateAdmin)
	g.POST("/teacher", CreateTeacherApi)
	g.GET("/teacher", FindTeacherApi)
	g.GET("/teachers", FindAllTeachersApi)
	g.PUT("/teacher", UpdateTeacherApi)
	g.DELETE("/teacher", DeleteTeacherApi)
	g.POST("/student", CreateStudentApi)
	g.GET("/student", FindStudentApi)
	g.GET("/students", FindAllStudentsApi)
	g.PUT("/student", UpdateStudentApi)
	g.DELETE("/student", DeleteStudentApi)
	g.POST("/event", CreateEventApi)
	g.GET("/event", FindEventApi)
	g.GET("/events", FindAllEventsApi)
	g.PUT("/event", UpdateEventApi)
	g.DELETE("/event", DeleteEventApi)
	g.POST("/registeredEvent", CreateRegisteredEventApi)
	g.GET("/registeredEvent/studentId", FindAllRegisteredEventByStudentIdApi)
	g.GET("/registeredEvent/eventId", FindAllRegisteredEventByEventIdApi)
	g.PUT("/registeredEvent", UpdateRegisteredEventApi)
	g.POST("/notice", CreateNoticeApi)
	g.GET("/notices", FindAllNoticesApi)
	g.PUT("/notice", UpdateNoticeApi)
	g.DELETE("/notice", DeleteNoticeApi)
	g.POST("/vote", CreateVoteApi)
	g.GET("/votes", FindAllVotesApi)
	g.PUT("/vote", UpdateVoteApi)
	g.DELETE("/vote", DeleteVoteApi)
	g.POST("/singleVote", CreateSingleVoteApi)
	g.PUT("/singleVote", UpdateSingleVoteApi)
	g.DELETE("/singleVote/voteId/studentId", DeleteSingleVoteApi)
	g.GET("/voteResults/voteId", GetVoteResultsApi)
}
