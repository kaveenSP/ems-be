package api

import (
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateRegisteredEventApi(c echo.Context) error {
	payloadObj := models.RegisteredEvent{}
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	if err := functions.UC_registeredEvent(payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.CreateRegisteredEvent(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindAllRegisteredEventByStudentIdApi(c echo.Context) error {
	studentId := c.QueryParam("studentId")
	returnVal, err := dbOperations.FindAllRegisteredEventByStudentId(studentId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindAllRegisteredEventByEventIdApi(c echo.Context) error {
	eventId := c.QueryParam("eventId")
	returnVal, err := dbOperations.FindAllRegisteredEventByEventId(eventId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func UpdateRegisteredEventApi(c echo.Context) error {
	payloadObj := models.RegisteredEvent{}
	if err := c.Bind(&payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.UpdateRegisteredEvent(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}
