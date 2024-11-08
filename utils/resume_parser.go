package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type ResumeResponse struct {
	Education []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"education"`
	Email      string `json:"email"`
	Experience []struct {
		Dates []string `json:"dates"`
		Name  string   `json:"name"`
		URL   string   `json:"url,omitempty"`
	} `json:"experience"`
	Name   string   `json:"name"`
	Phone  string   `json:"phone"`
	Skills []string `json:"skills"`
}

func ParseResume(filePath string) (ResumeResponse, error) {
	var response ResumeResponse

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", bytes.NewBuffer(fileData))
	if err != nil {
		return response, err
	}
	apiKey := os.Getenv("API_KEY")
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response, err
	}

	return response, nil
}
