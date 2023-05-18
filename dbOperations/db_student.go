package dbOperations

import (
	"context"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateStudent(student *models.Student) (*models.Student, error) {
	_, err := DB.Collection("Students").InsertOne(context.Background(), student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func FindStudent(studentId string) (*models.Student, error) {
	var student models.Student
	err := DB.Collection("Students").FindOne(context.Background(), bson.M{"studentid": studentId}).Decode(&student)
	if err != nil {
		return nil, err
	}
	if student == (models.Student{}) {
		return nil, errors.New("Teacher ID Not Found")
	}
	return &student, nil
}

func FindStudentByEmail(email string) (*models.Student, error) {
	var student models.Student
	err := DB.Collection("Students").FindOne(context.Background(), bson.M{"email": email}).Decode(&student)
	if err != nil {
		return nil, err
	}
	if student == (models.Student{}) {
		return nil, errors.New("Teacher ID Not Found")
	}
	return &student, nil
}

func FindAllStudents() (*[]models.Student, error) {
	var students []models.Student
	dbRes, err := DB.Collection("Students").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for dbRes.Next(context.Background()) {
		var student models.Student
		if err = dbRes.Decode(&student); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return &students, nil
}

func UpdateStudent(student *models.Student) (*models.Student, error) {
	_, err := DB.Collection("Students").UpdateOne(context.Background(), bson.M{"studentid": student.StudentId}, bson.M{"$set": student})
	if err != nil {
		return nil, err
	}
	return student, nil
}

func DeleteStudent(studentId string) (*models.Student, error) {
	res, err := DB.Collection("Students").DeleteOne(context.Background(), bson.M{"studentid": studentId})
	if err != nil {
		return nil, err
	}
	if res.DeletedCount < 1 {
		return nil, errors.New("Student ID Not Found")
	}
	return nil, nil
}
