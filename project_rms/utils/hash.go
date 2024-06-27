package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) ([]byte, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    return hashedPassword, nil
}

func CheckPasswordHash(password []byte, hash []byte) bool {
    err := bcrypt.CompareHashAndPassword(hash, password)
    return err == nil
}
