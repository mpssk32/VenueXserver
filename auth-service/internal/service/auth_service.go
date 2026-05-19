package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func Register(req models.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	return repository.CreateUser(user)
}

func Login(req models.LoginRequest) (string, error) {
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", err
	}

	token, err := GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}