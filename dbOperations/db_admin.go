package dbOperations

import (
	"context"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func FindAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	err := DB.Collection("Admins").FindOne(context.Background(), bson.M{"email": email}).Decode(&admin)
	if err != nil {
		return nil, err
	}
	if admin == (models.Admin{}) {
		return nil, errors.New("Teacher ID Not Found")
	}
	return &admin, nil
}
