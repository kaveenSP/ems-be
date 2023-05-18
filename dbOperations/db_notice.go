package dbOperations

import (
	"context"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateNotice(event *models.Notice) (*models.Notice, error) {
	_, err := DB.Collection("Notices").InsertOne(context.Background(), event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func FindAllTNotices() (*[]models.Notice, error) {
	var notices []models.Notice
	dbRes, err := DB.Collection("Notices").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for dbRes.Next(context.Background()) {
		var notice models.Notice
		if err = dbRes.Decode(&notice); err != nil {
			return nil, err
		}
		notices = append(notices, notice)
	}
	return &notices, nil
}

func UpdateNotice(notice *models.Notice) (*models.Notice, error) {
	_, err := DB.Collection("Notices").UpdateOne(context.Background(), bson.M{"noticeid": notice.NoticeId}, bson.M{"$set": notice})
	if err != nil {
		return nil, err
	}
	return notice, nil
}

func DeleteNotice(noticeId string) (*models.Notice, error) {
	res, err := DB.Collection("Notices").DeleteOne(context.Background(), bson.M{"noticeid": noticeId})
	if err != nil {
		return nil, err
	}
	if res.DeletedCount < 1 {
		return nil, errors.New("Notice ID Not Found")
	}
	return nil, nil
}
