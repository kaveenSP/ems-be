package api

import (
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTeacherApi(c echo.Context) error {
	payloadObj := models.Teacher{}
	if err := c.Bind(&payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := functions.UC_teacher(payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.CreateTeacher(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusBadRequest, returnVal)
}

func FindTeacherApi(c echo.Context) error {
	teacherId := c.QueryParam("teacherId")
	returnVal, err := dbOperations.FindTeacher(teacherId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusBadRequest, returnVal)
}

func FindAllTeacherApi(c echo.Context) error {
	returnVal, err := dbOperations.FindAllTeachers()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusBadRequest, returnVal)
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
	return c.JSON(http.StatusBadRequest, returnVal)
}

func DeleteTeacherApi(c echo.Context) error {
	teacherId := c.QueryParam("teacherId")
	returnVal, err := dbOperations.DeleteTeacher(teacherId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusBadRequest, returnVal)
}
