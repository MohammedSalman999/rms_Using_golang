package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/MohammedSalman999/rms.go/config"
	middlewares "github.com/MohammedSalman999/rms.go/middleware"
	"github.com/MohammedSalman999/rms.go/models"
	"github.com/MohammedSalman999/rms.go/utils"
)

func UploadResume(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("resume")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	email := r.Context().Value(middlewares.ContextKeyEmail).(string)
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filePath := fmt.Sprintf("./uploads/%s", handler.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resumeData, err := utils.ParseResume(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile := models.Profile{
		UserID:     user.ID,
		ResumeFile: filePath,
		Skills:     strings.Join(resumeData.Skills, ", "),
		Education:  resumeData.Education[0].Name,
		Experience: resumeData.Experience[0].Name,
		Name:       resumeData.Name,
		Phone:      resumeData.Phone,
	}

	if err := config.DB.Save(&profile).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
