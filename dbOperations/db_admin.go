package dbOperations

import (
	"context"
	"ems-be/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateAdmin(event *models.Admin) (*models.Admin, error) {
	_, err := DB.Collection("Admins").InsertOne(context.Background(), event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func FindAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	err := DB.Collection("Admins").FindOne(context.Background(), bson.M{"email": email}).Decode(&admin)
	if err != nil {
		return nil, err
	}
	if admin == (models.Admin{}) {
		return nil, errors.New("Admin ID Not Found")
	}
	return &admin, nil
}

func UpdateAdmin(admin *models.Admin) (*models.Admin, error) {
	_, err := DB.Collection("Admins").UpdateOne(context.Background(), bson.M{"adminid": admin.AdminId}, bson.M{"$set": admin})
	if err != nil {
		return nil, err
	}
	return admin, nil
}
