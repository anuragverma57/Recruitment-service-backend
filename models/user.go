package models

type User struct {
	ID              string `json:"id" bson:"_id,omitempty"`
	Name            string `json:"name" bson:"name"`
	Email           string `json:"email" bson:"email"`
	Address         string `json:"address" bson:"address"`
	UserType        string `json:"user_type" bson:"user_type"` // Admin or Applicant
	PasswordHash    string `json:"password_hash" bson:"password_hash"`
	ProfileHeadline string `json:"profile_headline" bson:"profile_headline"`
}
