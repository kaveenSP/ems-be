package api

import (
	"ems-be/amazon_s3"
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateEventApi(c echo.Context) error {
	payloadObj := models.Event{}
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	if err := functions.UC_event(payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Create an AWS session and S3 client
	svc, err := amazon_s3.InitiateConnectionWithImageService()
	print(svc)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Upload the image to S3
	imagePath, err := amazon_s3.UploadImageToS3(svc, payloadObj.ImagePath)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Set the S3 image path in the event object
	payloadObj.ImagePath = imagePath

	// Save the event to the database
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
	_, err = dbOperations.DeleteAllRegisteredEventsByEventId(eventId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}
