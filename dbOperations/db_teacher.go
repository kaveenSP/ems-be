package dbOperations

import (
	"context"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTeacher(teacher *models.Teacher) (*models.Teacher, error) {
	_, err := DB.Collection("Teachers").InsertOne(context.Background(), teacher)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func FindTeacher(teacherId string) (*models.Teacher, error) {
	var teacher models.Teacher
	err := DB.Collection("Teachers").FindOne(context.Background(), bson.M{"teacherid": teacherId}).Decode(&teacher)
	if err != nil {
		return nil, err
	}
	if teacher == (models.Teacher{}) {
		return nil, errors.New("Teacher ID Not Found")
	}
	return &teacher, nil
}

func FindTeacherByEmail(email string) (*models.Teacher, error) {
	var teacher models.Teacher
	err := DB.Collection("Teachers").FindOne(context.Background(), bson.M{"email": email}).Decode(&teacher)
	if err != nil {
		return nil, err
	}
	if teacher == (models.Teacher{}) {
		return nil, errors.New("Teacher ID Not Found")
	}
	return &teacher, nil
}

func FindAllTeachers() (*[]models.Teacher, error) {
	var teachers []models.Teacher
	dbRes, err := DB.Collection("Teachers").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for dbRes.Next(context.Background()) {
		var teacher models.Teacher
		if err = dbRes.Decode(&teacher); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return &teachers, nil
}

func UpdateTeacher(teacher *models.Teacher) (*models.Teacher, error) {
	_, err := DB.Collection("Teachers").UpdateOne(context.Background(), bson.M{"teacherid": teacher.TeacherId}, bson.M{"$set": teacher})
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func DeleteTeacher(teacherId string) (*models.Teacher, error) {
	res, err := DB.Collection("Teachers").DeleteOne(context.Background(), bson.M{"teacherid": teacherId})
	if err != nil {
		return nil, err
	}
	if res.DeletedCount < 1 {
		return nil, errors.New("Teacher ID Not Found")
	}
	return nil, nil
}
