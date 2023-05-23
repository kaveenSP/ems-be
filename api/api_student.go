package api

import (
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateStudentApi(c echo.Context) error {
	payloadObj := models.Student{}
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	if err := functions.UC_student(payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.CreateStudent(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindStudentApi(c echo.Context) error {
	studentId := c.QueryParam("studentId")
	returnVal, err := dbOperations.FindStudent(studentId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindAllStudentsApi(c echo.Context) error {
	returnVal, err := dbOperations.FindAllStudents()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func UpdateStudentApi(c echo.Context) error {
	payloadObj := models.Student{}
	if err := c.Bind(&payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.UpdateStudent(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func DeleteStudentApi(c echo.Context) error {
	studentId := c.QueryParam("studentId")
	returnVal, err := dbOperations.DeleteStudent(studentId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}
