package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MohammedSalman999/rms.go/config"
	"github.com/MohammedSalman999/rms.go/models"
	"github.com/MohammedSalman999/rms.go/utils"
	"gorm.io/gorm"
)

type AuthResponse struct {
    Token string `json:"token"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedPassword, err := utils.HashPassword([]byte(user.PasswordHash))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user.PasswordHash = hashedPassword

    if err := config.DB.Create(&user).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Email    string `json:"email"`
        Password []byte `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(user.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(AuthResponse{Token: token})
}
