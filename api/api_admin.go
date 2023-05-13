package api

import (
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateAdminApi(c echo.Context) error {
	payloadObj := models.Admin{}
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	if err := functions.UC_admin(payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.CreateAdmin(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal.Password = ""
	return c.JSON(http.StatusOK, returnVal)
}

func UpdateAdmin(c echo.Context) error {
	payloadObj := models.Admin{}
	if err := c.Bind(&payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.UpdateAdmin(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}
