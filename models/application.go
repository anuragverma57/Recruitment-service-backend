package models

type Application struct {
	UserID string `json:"user_id" bson:"user_id"`
	JobID  string `json:"job_id" bson:"job_id"`
}
