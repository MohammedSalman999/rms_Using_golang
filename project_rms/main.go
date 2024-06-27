package main

import (
	
	"log"
	"net/http"

	"github.com/MohammedSalman999/rms.go/config"
	"github.com/MohammedSalman999/rms.go/handlers"
	"github.com/MohammedSalman999/rms.go/middleware"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	

	r := mux.NewRouter()

	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)
	api.HandleFunc("/uploadResume", handlers.UploadResume).Methods("POST")
	api.HandleFunc("/admin/job", handlers.CreateJob).Methods("POST")
	api.HandleFunc("/admin/job/{job_id}", handlers.GetJob).Methods("GET")
	api.HandleFunc("/admin/applicants", handlers.GetApplicants).Methods("GET")
	api.HandleFunc("/admin/applicant/{applicant_id}", handlers.GetApplicant).Methods("GET")
	api.HandleFunc("/jobs", handlers.GetJobs).Methods("GET")
	api.HandleFunc("/jobs/apply", handlers.ApplyJob).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
