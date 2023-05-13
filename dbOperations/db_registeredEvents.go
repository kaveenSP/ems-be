package dbOperations

import (
	"context"
	"ems-be/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateRegisteredEvent(event *models.RegisteredEvent) (*models.RegisteredEvent, error) {
	_, err := DB.Collection("RegisteredEvents").InsertOne(context.Background(), event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func FindAllRegisteredEventByStudentId(studentId string) (*[]models.RegisteredEvent, error) {
	var registeredEvents []models.RegisteredEvent
	dbRes, err := DB.Collection("RegisteredEvents").Find(context.Background(), bson.M{"studentid": studentId})
	if err != nil {
		return nil, err
	}
	for dbRes.Next(context.Background()) {
		var registeredEvent models.RegisteredEvent
		if err = dbRes.Decode(&registeredEvent); err != nil {
			return nil, err
		}
		registeredEvents = append(registeredEvents, registeredEvent)
	}
	return &registeredEvents, nil
}

func FindAllRegisteredEventByEventId(eventId string) (*[]models.RegisteredEvent, error) {
	var registeredEvents []models.RegisteredEvent
	dbRes, err := DB.Collection("RegisteredEvents").Find(context.Background(), bson.M{"eventid": eventId})
	if err != nil {
		return nil, err
	}
	for dbRes.Next(context.Background()) {
		var registeredEvent models.RegisteredEvent
		if err = dbRes.Decode(&registeredEvent); err != nil {
			return nil, err
		}
		registeredEvents = append(registeredEvents, registeredEvent)
	}
	return &registeredEvents, nil
}

func UpdateRegisteredEvent(registeredEvent *models.RegisteredEvent) (*models.RegisteredEvent, error) {
	_, err := DB.Collection("RegisteredEvents").UpdateOne(context.Background(), bson.M{"studentid": registeredEvent.StudentId}, bson.M{"$set": registeredEvent})
	if err != nil {
		return nil, err
	}
	return registeredEvent, nil
}

func DeleteAllRegisteredEventsByEventId(eventId string) (*models.RegisteredEvent, error) {
	_, err := DB.Collection("Teachers").DeleteMany(context.Background(), bson.M{"eventid": eventId})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
