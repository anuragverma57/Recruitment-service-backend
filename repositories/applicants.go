package repositories

import (
	"context"
	"recruitment-system/config"
	"recruitment-system/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var applicantCollection *mongo.Collection

// GetAllApplicants retrieves all applicants from the database
func GetAllApplicants(ctx context.Context) ([]models.User, error) {
	applicantCollection = config.MongoClient.Database("recruitment_system").Collection("applicants")
	var applicants []models.User
	cursor, err := applicantCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var applicant models.User
		if err := cursor.Decode(&applicant); err != nil {
			return nil, err
		}
		applicants = append(applicants, applicant)
	}
	return applicants, nil
}

// GetApplicantByID retrieves a specific applicant by their ID
func GetApplicantByID(ctx context.Context, id string) (models.User, error) {
	var applicant models.User
	err := applicantCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&applicant)
	return applicant, err
}
