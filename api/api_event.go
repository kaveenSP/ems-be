package api

import (
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateEventApi(c echo.Context) error {
	payloadObj := models.Event{}
	//bind payload data
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	//unique check
	if err := functions.UC_event(payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.CreateEvent(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindEventApi(c echo.Context) error {
	eventId := c.QueryParam("eventId")
	returnVal, err := dbOperations.FindEvent(eventId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindAllEventsApi(c echo.Context) error {
	returnVal, err := dbOperations.FindAllTEvents()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func UpdateEventApi(c echo.Context) error {
	payloadObj := models.Event{}
	if err := c.Bind(&payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.UpdateEvent(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func DeleteEventApi(c echo.Context) error {
	eventId := c.QueryParam("eventId")
	returnVal, err := dbOperations.DeleteEvent(eventId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}
