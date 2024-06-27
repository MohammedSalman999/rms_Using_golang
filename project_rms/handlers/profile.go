package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MohammedSalman999/rms.go/config"
	"github.com/MohammedSalman999/rms.go/middleware" // Import middleware package
	"github.com/MohammedSalman999/rms.go/models"
	"github.com/gorilla/mux"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []models.Job
	if err := config.DB.Find(&jobs).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(jobs)
}

func ApplyJob(w http.ResponseWriter, r *http.Request) {
	// Get job_id from URL query parameters
	jobIDStr := mux.Vars(r)["job_id"]
	if jobIDStr == "" {
		http.Error(w, "Job ID is required", http.StatusBadRequest)
		return
	}

	// Convert jobIDStr to uint
	jobID, err := strconv.ParseUint(jobIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Job ID", http.StatusBadRequest)
		return
	}

	// Extract email from context (assuming it's set by your authentication middleware)
	email := r.Context().Value(middlewares.ContextKeyEmail).(string)

	// Retrieve user based on email
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create application record
	// Assuming JobID in models.Application is of type string
	application := models.Application{
		JobID:    uint(jobID),
		Applicant: user.ID, // Assuming user.ID is of type uint
	}

	// Save application to database
	if err := config.DB.Create(&application).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with HTTP status created (201)
	w.WriteHeader(http.StatusCreated)
}
