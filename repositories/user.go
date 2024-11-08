package repositories

import (
	"context"
	"fmt"
	"recruitment-system/config"
	"recruitment-system/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func CreateUser(ctx context.Context, user models.User) error {
	userCollection = config.MongoClient.Database("recruitment_system").Collection("users")
	_, err := userCollection.InsertOne(ctx, user)
	return err
}

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	userCollection = config.MongoClient.Database("recruitment_system").Collection("users")
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}

func GetUserByID(ctx context.Context, id string) models.User {
	var user models.User
	userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	fmt.Println(user)
	return user
}
