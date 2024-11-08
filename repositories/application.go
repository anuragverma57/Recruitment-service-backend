package repositories

import (
	"context"
	"recruitment-system/config"
	"recruitment-system/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var applicationCollection *mongo.Collection

// ApplyForJob saves a job application in the database
func ApplyForJob(ctx context.Context, application models.Application) error {
	applicationCollection = config.MongoClient.Database("recruitment_system").Collection("applications")

	_, err := applicationCollection.InsertOne(ctx, application)
	return err
}
