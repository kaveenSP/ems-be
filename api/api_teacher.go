package api

import (
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTeacherApi(c echo.Context) error {
	payloadObj := models.Teacher{}
	//bind payload data
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	//unique check
	if err := functions.UC_teacher(payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.CreateTeacher(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal.Password = ""
	return c.JSON(http.StatusOK, returnVal)
}

func FindTeacherApi(c echo.Context) error {
	teacherId := c.QueryParam("teacherId")
	returnVal, err := dbOperations.FindTeacher(teacherId)
	fmt.Println(returnVal)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindAllTeacherApi(c echo.Context) error {
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
