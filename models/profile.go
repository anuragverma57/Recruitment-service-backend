package models

type Profile struct {
	Applicant         User   `json:"applicant" bson:"applicant"`
	ResumeFileAddress string `json:"resume_file_address" bson:"resume_file_address"`
	Skills            string `json:"skills" bson:"skills"`
	Education         string `json:"education" bson:"education"`
	Experience        string `json:"experience" bson:"experience"`
}
