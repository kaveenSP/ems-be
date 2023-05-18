package dbOperations

import (
	"context"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateVote(event *models.Vote) (*models.Vote, error) {
	_, err := DB.Collection("Votes").InsertOne(context.Background(), event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func FindVote(voteId string) (*models.Vote, error) {
	var vote models.Vote
	err := DB.Collection("Votes").FindOne(context.Background(), bson.M{"voteid": voteId}).Decode(&vote)
	if err != nil {
		return nil, err
	}
	return &vote, nil
}

func FindAllTVotes() (*[]models.Vote, error) {
	var votes []models.Vote
	dbRes, err := DB.Collection("Votes").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for dbRes.Next(context.Background()) {
		var vote models.Vote
		if err = dbRes.Decode(&vote); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return &votes, nil
}

func UpdateVote(vote *models.Vote) (*models.Vote, error) {
	_, err := DB.Collection("Votes").UpdateOne(context.Background(), bson.M{"voteid": vote.VoteId}, bson.M{"$set": vote})
	if err != nil {
		return nil, err
	}
	return vote, nil
}

func DeleteVote(voteId string) (*models.Vote, error) {
	res, err := DB.Collection("Votes").DeleteOne(context.Background(), bson.M{"voteid": voteId})
	if err != nil {
		return nil, err
	}
	if res.DeletedCount < 1 {
		return nil, errors.New("Vote ID Not Found")
	}
	return nil, nil
}
