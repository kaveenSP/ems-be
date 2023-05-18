package dbOperations

import (
	"context"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateSingleVote(event *models.SingleVote) (*models.SingleVote, error) {
	_, err := DB.Collection("SingleVotes").InsertOne(context.Background(), event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func FindAllTSingleVotesByVoteId(voteId string) (*[]models.SingleVote, error) {
	var singleVotes []models.SingleVote
	dbRes, err := DB.Collection("SingleVotes").Find(context.Background(), bson.M{"voteid": voteId})
	if err != nil {
		return nil, err
	}
	for dbRes.Next(context.Background()) {
		var singleVote models.SingleVote
		if err = dbRes.Decode(&singleVote); err != nil {
			return nil, err
		}
		singleVotes = append(singleVotes, singleVote)
	}
	return &singleVotes, nil
}

func FindAllTSingleVotesByVoteIdAndOption(voteId string, option string) (*[]models.SingleVote, error) {
	var singleVotes []models.SingleVote
	dbRes, err := DB.Collection("SingleVotes").Find(context.Background(), bson.M{"voteid": voteId, "option": option})
	if err != nil {
		return nil, err
	}
	for dbRes.Next(context.Background()) {
		var singleVote models.SingleVote
		if err = dbRes.Decode(&singleVote); err != nil {
			return nil, err
		}
		singleVotes = append(singleVotes, singleVote)
	}
	return &singleVotes, nil
}

func UpdateSingleVote(singleVote *models.SingleVote) (*models.SingleVote, error) {
	_, err := DB.Collection("SingleVotes").UpdateOne(context.Background(), bson.M{"studentid": singleVote.StudentId, "voteid": singleVote.VoteId}, bson.M{"$set": singleVote})
	if err != nil {
		return nil, err
	}
	return singleVote, nil
}

func DeleteSingleVote(voteId string, studentId string) (*models.Student, error) {
	res, err := DB.Collection("SingleVotes").DeleteOne(context.Background(), bson.M{"voteId": voteId, "studentid": studentId})
	if err != nil {
		return nil, err
	}
	if res.DeletedCount < 1 {
		return nil, errors.New("Student ID Not Found")
	}
	return nil, nil
}

func DeleteAllSingleVotesByVoteId(voteId string) (*models.SingleVote, error) {
	_, err := DB.Collection("SingleVotes").DeleteMany(context.Background(), bson.M{"voteid": voteId})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
