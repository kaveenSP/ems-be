package api

import (
	"ems-be/dbOperations"
	"ems-be/models"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func LoginApi(c echo.Context) error {
	payloadObj := models.Login{}
	var admin *models.Admin
	var teacher *models.Teacher
	var student *models.Student
	var err error
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	if strings.EqualFold(payloadObj.Role, "admin") {
		admin, err = dbOperations.FindAdminByEmail(payloadObj.Email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, errors.New("User Not Found"))
		}
		if admin.Password != payloadObj.Password {
			return c.JSON(http.StatusUnauthorized, errors.New("Incorrect Password"))
		}
		return c.JSON(http.StatusOK, admin)
	} else if strings.EqualFold(payloadObj.Role, "teacher") {
		teacher, err = dbOperations.FindTeacherByEmail(payloadObj.Email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, errors.New("User Not Found"))
		}
		if teacher.Password != payloadObj.Password {
			return c.JSON(http.StatusUnauthorized, errors.New("Incorrect Password"))
		}
		return c.JSON(http.StatusOK, teacher)
	} else if strings.EqualFold(payloadObj.Role, "student") {
		student, err = dbOperations.FindStudentByEmail(payloadObj.Email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, errors.New("User Not Found"))
		}
		if student.Password != payloadObj.Password {
			return c.JSON(http.StatusUnauthorized, errors.New("Incorrect Password"))
		}
		return c.JSON(http.StatusOK, student)
	}
	return c.JSON(http.StatusBadRequest, errors.New("Unidentified Role"))
}
