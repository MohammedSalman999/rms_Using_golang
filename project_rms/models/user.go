package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string `json:"name"`
	Email           string `json:"email" gorm:"unique"`
	PasswordHash    []byte `json:"-"`
	UserType        string `json:"user_type"`
	ProfileHeadline string `json:"profile_headline"`
	Address         string `json:"address"`
	Profile         Profile `gorm:"foreignKey:UserID"`
}

type Profile struct {
	gorm.Model
	UserID         uint   `json:"-"`
	ResumeFile     string `json:"resume_file"`
	Skills         string `json:"skills"`
	Education      string `json:"education"`
	Experience     string `json:"experience"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
}
