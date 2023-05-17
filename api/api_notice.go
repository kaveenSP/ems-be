package api

import (
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateNoticeApi(c echo.Context) error {
	payloadObj := models.Notice{}
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	if err := functions.UC_notice(payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.CreateNotice(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindAllNoticesApi(c echo.Context) error {
	returnVal, err := dbOperations.FindAllTNotices()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func UpdateNoticeApi(c echo.Context) error {
	payloadObj := models.Notice{}
	if err := c.Bind(&payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.UpdateNotice(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func DeleteNoticeApi(c echo.Context) error {
	noticeId := c.QueryParam("noticeId")
	returnVal, err := dbOperations.DeleteNotice(noticeId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}
