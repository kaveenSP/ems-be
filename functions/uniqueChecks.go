package functions

import (
	"context"
	"ems-be/dbOperations"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func UC_admin(object models.Admin) error {
	admin := models.Admin{}
	err := dbOperations.DB.Collection("Admins").FindOne(context.Background(), bson.M{"adminid": object.AdminId}).Decode(&admin)
	if err != nil {
		return nil
	}
	return errors.New(admin.FirstName + " " + "Already Exists")
}

func UC_teacher(object models.Teacher) error {
	teacher := models.Teacher{}
	err := dbOperations.DB.Collection("Teachers").FindOne(context.Background(), bson.M{"teacherid": object.TeacherId}).Decode(&teacher)
	if err != nil {
		return nil
	}
	return errors.New(teacher.FirstName + " " + "Already Exists")
}

func UC_student(object models.Student) error {
	teacher := models.Teacher{}
	err := dbOperations.DB.Collection("Students").FindOne(context.Background(), bson.M{"studentid": object.StudentId}).Decode(&teacher)
	if err != nil {
		return nil
	}
	return errors.New(teacher.FirstName + " " + "Already Exists")
}

func UC_event(object models.Event) error {
	event := models.Event{}
	err := dbOperations.DB.Collection("Events").FindOne(context.Background(), bson.M{"eventid": object.EventId}).Decode(&event)
	if err != nil {
		return nil
	}
	return errors.New(event.Name + " " + "Already Exists")
}

func UC_registeredEvent(object models.RegisteredEvent) error {
	event := models.Event{}
	err := dbOperations.DB.Collection("RegisteredEvents").FindOne(context.Background(), bson.M{"eventid": object.EventId, "studentid": object.StudentId}).Decode(&event)
	if err != nil {
		return nil
	}
	return errors.New(event.Name + " " + "Already Exists")
}

func UC_notice(object models.Notice) error {
	notice := models.Notice{}
	err := dbOperations.DB.Collection("Notices").FindOne(context.Background(), bson.M{"noticeId": object.NoticeId}).Decode(&notice)
	if err != nil {
		return nil
	}
	return errors.New(notice.Subject + " " + "Already Exists")
}
