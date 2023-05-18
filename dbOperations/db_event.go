package dbOperations

import (
	"context"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateEvent(event *models.Event) (*models.Event, error) {
	_, err := DB.Collection("Events").InsertOne(context.Background(), event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func FindEvent(eventId string) (*models.Event, error) {
	var event models.Event
	err := DB.Collection("Events").FindOne(context.Background(), bson.M{"eventid": eventId}).Decode(&event)
	if err != nil {
		return nil, err
	}
	if event == (models.Event{}) {
		return nil, errors.New("Event ID Not Found")
	}
	return &event, nil
}

func FindAllTEvents() (*[]models.Event, error) {
	var events []models.Event
	dbRes, err := DB.Collection("Events").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for dbRes.Next(context.Background()) {
		var event models.Event
		if err = dbRes.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return &events, nil
}

func UpdateEvent(event *models.Event) (*models.Event, error) {
	_, err := DB.Collection("Events").UpdateOne(context.Background(), bson.M{"eventid": event.EventId}, bson.M{"$set": event})
	if err != nil {
		return nil, err
	}
	return event, nil
}

func DeleteEvent(eventId string) (*models.Event, error) {
	res, err := DB.Collection("Events").DeleteOne(context.Background(), bson.M{"eventid": eventId})
	if err != nil {
		return nil, err
	}
	if res.DeletedCount < 1 {
		return nil, errors.New("Teacher ID Not Found")
	}
	return nil, nil
}
