package functions

import (
	"context"
	"ems-be/dbOperations"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func UC_teacher(object models.Teacher) error {
	teacher := models.Teacher{}
	err := dbOperations.DB.Collection("Teachers").FindOne(context.Background(), bson.M{"teacherid": teacher.TeacherId}).Decode(&teacher)
	if err != nil {

		return nil
	}
	return errors.New(object.FirstName + " " + "Already Exists")
}
