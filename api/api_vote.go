package api

import (
	"ems-be/dbOperations"
	"ems-be/functions"
	"ems-be/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func CreateVoteApi(c echo.Context) error {
	payloadObj := models.Vote{}
	if err0 := c.Bind(&payloadObj); err0 != nil {
		return c.String(http.StatusBadRequest, err0.Error())
	}
	if err := functions.UC_vote(payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Convert the date into the desired format
	payloadObj.CreatedAt = time.Now().Format("02/01/2006")
	returnVal, err := dbOperations.CreateVote(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func FindAllVotesApi(c echo.Context) error {
	returnVal, err := dbOperations.FindAllTVotes()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func UpdateVoteApi(c echo.Context) error {
	payloadObj := models.Vote{}
	if err := c.Bind(&payloadObj); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	returnVal, err := dbOperations.UpdateVote(&payloadObj)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func DeleteVoteApi(c echo.Context) error {
	voteId := c.QueryParam("voteId")
	returnVal, err := dbOperations.DeleteVote(voteId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	_, err = dbOperations.DeleteAllSingleVotesByVoteId(voteId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, returnVal)
}

func GetVoteResultsApi(c echo.Context) error {
	var voteResults []models.VoteResult

	voteId := c.QueryParam("voteId")
	vote, err := dbOperations.FindVote(voteId)
	if err != nil {
		return err
	}
	singleVotes, err := dbOperations.FindAllTSingleVotesByVoteId(voteId)

	if err != nil {
		return err
	}
	for _, option := range vote.Options {
		sVoteO, err := dbOperations.FindAllTSingleVotesByVoteIdAndOption(voteId, option)
		if err != nil {
			return err
		}
		voteResult := models.VoteResult{
			Option:     option,
			Percentage: fmt.Sprintf("%.2f", float64(len(*sVoteO))/float64(len(*singleVotes))*100),
		}
		voteResults = append(voteResults, voteResult)
	}
	return c.JSON(http.StatusOK, voteResults)
}
