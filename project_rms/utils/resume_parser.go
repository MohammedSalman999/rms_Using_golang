package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// here is the resume data struct
// which contains the values as per requirements

type ResumeData struct {
	Education []struct {
		Name string `json:"name"`
	} `json:"education"`
	Email      string `json:"email"`
	Experience []struct {
		Name string `json:"name"`
	} `json:"experience"`
	Name   string   `json:"name"`
	Phone  string   `json:"phone"`
	Skills []string `json:"skills"`
}
// parse resume will take the pathpath and will return the 
// resume data feilds or an error

func ParseResume(filePath string) (*ResumeData, error) {
	apiKey := "Apni_api_Dalo"
	url := "https://api.apilayer.com/resume_parser/upload"

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(fileBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var resumeData ResumeData
	err = json.NewDecoder(resp.Body).Decode(&resumeData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &resumeData, nil
}
