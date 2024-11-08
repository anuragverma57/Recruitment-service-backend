package repositories

import (
	"context"
	"recruitment-system/config"
	"recruitment-system/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var jobCollection *mongo.Collection

func CreateJob(ctx context.Context, job models.Job) error {
	jobCollection = config.MongoClient.Database("recruitment_system").Collection("jobs")

	_, err := jobCollection.InsertOne(ctx, job)
	return err
}

func GetAllJobs(ctx context.Context) ([]models.Job, error) {
	jobCollection = config.MongoClient.Database("recruitment_system").Collection("jobs")

	var jobs []models.Job
	cursor, err := jobCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var job models.Job
		if err := cursor.Decode(&job); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func GetJobByID(ctx context.Context, jobID string) (models.Job, error) {
	jobCollection = config.MongoClient.Database("recruitment_system").Collection("jobs")

	var job models.Job
	err := jobCollection.FindOne(ctx, bson.M{"_id": jobID}).Decode(&job)
	return job, err
}
