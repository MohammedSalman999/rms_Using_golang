package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/MohammedSalman999/rms.go/config"
	"github.com/MohammedSalman999/rms.go/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	job.PostedOn = time.Now()

	if err := config.DB.Create(&job).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobID, err := strconv.Atoi(vars["job_id"])
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	var job models.Job
	if err := config.DB.Preload("Applications").First(&job, jobID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Job not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(job)
}

func GetApplicants(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := config.DB.Where("user_type = ?", "applicant").Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetApplicant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	applicantID, err := strconv.Atoi(vars["applicant_id"])
	if err != nil {
		http.Error(w, "Invalid applicant ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := config.DB.Preload("Profile").First(&user, applicantID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Applicant not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
