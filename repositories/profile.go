package repositories

import (
	"context"
	"recruitment-system/config"
	"recruitment-system/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var profileCollection *mongo.Collection

// SaveProfile saves the applicant's profile in the database
func SaveProfile(ctx context.Context, profile models.Profile) error {
	profileCollection = config.MongoClient.Database("recruitment_system").Collection("profiles")
	_, err := profileCollection.InsertOne(ctx, profile)
	return err
}
