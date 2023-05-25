package api

import (
	"ems-be/amazon_s3"
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTeacherApi(c echo.Context) error {
	payloadObj := models.Teacher{}
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	if err := functions.UC_teacher(payloadObj); err != nil {
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
	returnVal, err := dbOperations.CreateTeacher(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, returnVal)
}

func FindTeacherApi(c echo.Context) error {
	teacherId := c.QueryParam("teacherId")
	returnVal, err := dbOperations.FindTeacher(teacherId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindAllTeachersApi(c echo.Context) error {
	returnVal, err := dbOperations.FindAllTeachers()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func UpdateTeacherApi(c echo.Context) error {
	payloadObj := models.Teacher{}
	if err := c.Bind(&payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.UpdateTeacher(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func DeleteTeacherApi(c echo.Context) error {
	teacherId := c.QueryParam("teacherId")
	returnVal, err := dbOperations.DeleteTeacher(teacherId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}
